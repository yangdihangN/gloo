package conversion

import (
	"context"

	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"go.uber.org/zap"
)

var (
	FailedToListGatewayResourcesError = func(err error, version, namespace string) error {
		return errors.Wrapf(err, "Failed to list %v gateway resources in %v", version, namespace)
	}

	FailedToConvertGatewayError = func(err error, version, namespace, name string) error {
		return errors.Wrapf(err, "Failed to convert %v gateway %v.%v", version, namespace, name)
	}

	FailedToDeleteGatewayError = func(err error, version, namespace, name string) error {
		return errors.Wrapf(err, "Failed to delete %v gateway %v.%v", version, namespace, name)
	}

	FailedToWriteGatewayError = func(err error, version, namespace, name string) error {
		return errors.Wrapf(err, "Failed to write %v gateway %v.%v", version, namespace, name)
	}
)

type Ladder interface {
	Climb()
}

// TODO use solo-kit's interface
type Converter interface {
	Convert(src, dst resources.Resource) error
}

type ladder struct {
	ctx               context.Context
	namespace         string
	v1Client          v1.GatewayClient
	v2alpha1Client    v2alpha1.GatewayClient
	v2alpha1Converter Converter
}

func NewLadder(
	ctx context.Context,
	namespace string,
	v1Client v1.GatewayClient,
	v2alpha1Client v2alpha1.GatewayClient,
	gatewayConverter Converter,
) Ladder {

	return &ladder{
		ctx:               ctx,
		namespace:         namespace,
		v1Client:          v1Client,
		v2alpha1Client:    v2alpha1Client,
		v2alpha1Converter: gatewayConverter,
	}
}

// With more rungs we could (read, convert, read & merge, convert, ...,  write)
func (c *ladder) Climb() {
	v1List, err := c.v1Client.List(c.namespace, clients.ListOpts{Ctx: c.ctx})
	if err != nil {
		wrapped := FailedToListGatewayResourcesError(err, "v1", c.namespace)
		contextutils.LoggerFrom(c.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.String("namespace", c.namespace))
	}

	v2alpha1List := make([]*v2alpha1.Gateway, 0, len(v1List))
	for _, oldGateway := range v1List {
		convertedGateway := &v2alpha1.Gateway{}

		err := c.v2alpha1Converter.Convert(oldGateway, convertedGateway)
		if err != nil {
			wrapped := FailedToConvertGatewayError(
				err,
				"v1",
				oldGateway.GetMetadata().GetNamespace(),
				oldGateway.GetMetadata().GetName())
			contextutils.LoggerFrom(c.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("gateway", oldGateway))
		} else {
			v2alpha1List = append(v2alpha1List, convertedGateway)
		}

		if err := c.v1Client.Delete(
			convertedGateway.GetMetadata().GetNamespace(),
			convertedGateway.GetMetadata().GetName(),
			clients.DeleteOpts{Ctx: c.ctx}); err != nil {

			wrapped := FailedToDeleteGatewayError(
				err,
				"v1",
				convertedGateway.GetMetadata().GetNamespace(),
				convertedGateway.GetMetadata().GetName())
			contextutils.LoggerFrom(c.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("gateway", convertedGateway))
		}
	}

	for _, newGateway := range v2alpha1List {
		_, err := c.v2alpha1Client.Write(newGateway, clients.WriteOpts{Ctx: c.ctx})
		if err != nil {
			wrapped := FailedToWriteGatewayError(
				err,
				"v2alpha1",
				newGateway.GetMetadata().GetNamespace(),
				newGateway.GetMetadata().GetName())
			contextutils.LoggerFrom(c.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("gateway", newGateway))
		}
	}
}
