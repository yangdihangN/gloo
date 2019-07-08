package convertgateway

import (
	"context"

	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"go.uber.org/zap"
)

var (
	FailedToListGatewayResources = func(err error, version, namespace string) error {
		return errors.Wrapf(err, "Failed to list %v gateway resources in %v", version, namespace)
	}

	FailedToDeleteGateway = func(err error, version, namespace, name string) error {
		return errors.Wrapf(err, "Failed to delete %v gateway %v.%v", version, namespace)
	}

	FailedToWriteGateway = func(err error, version, namespace, name string) error {
		return errors.Wrapf(err, "Failed to write %v gateway %v.%v", version, namespace)
	}
)

// TODO use solo-kit's interface
type Ladder interface {
	Climb()
}

type ladder struct {
	ctx               context.Context
	namespace         string
	v1Client          v1.GatewayClient
	v2alpha1Client    v2alpha1.GatewayClient
	v2alpha1Converter V2alpha1Converter
}

func NewLadder(
	ctx context.Context,
	namespace string,
	v1Client v1.GatewayClient,
	v2alpha1Client v2alpha1.GatewayClient,
	gatewayConverter V2alpha1Converter,
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
		wrapped := FailedToListGatewayResources(err, "v1", c.namespace)
		contextutils.LoggerFrom(c.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.String("namespace", c.namespace))
	}

	v2alpha1List := make([]*v2alpha1.Gateway, len(v1List), len(v1List))
	for _, oldGateway := range v1List {
		convertedGateway := c.v2alpha1Converter.Convert(oldGateway)
		v2alpha1List = append(v2alpha1List, convertedGateway)

		if err := c.v1Client.Delete(
			convertedGateway.GetMetadata().GetNamespace(),
			convertedGateway.GetMetadata().GetName(),
			clients.DeleteOpts{Ctx: c.ctx}); err != nil {

			wrapped := FailedToDeleteGateway(
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
			wrapped := FailedToWriteGateway(
				err,
				"v2alpha1",
				newGateway.GetMetadata().GetNamespace(),
				newGateway.GetMetadata().GetName())
			contextutils.LoggerFrom(c.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("gateway", newGateway))
		}
	}
}
