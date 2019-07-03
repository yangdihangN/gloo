package controller

import (
	"context"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	"github.com/solo-io/gloo/projects/gatewayinit/pkg/conversion"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"go.uber.org/zap"
)

type Controller interface {
	DoConversion()
}

type controller struct {
	ctx            context.Context
	v1Client       v1.GatewayClient
	v2alpha1Client v2alpha1.GatewayClient
	converter      conversion.GatewayConverter
	namespace      string
}

func NewController(
	ctx context.Context,
	v1Client v1.GatewayClient,
	v2alpha1Client v2alpha1.GatewayClient,
	converter conversion.GatewayConverter,
	namespace string,
) Controller {

	return &controller{
		ctx:            ctx,
		v1Client:       v1Client,
		v2alpha1Client: v2alpha1Client,
		converter:      converter,
		namespace:      namespace,
	}
}

func (c *controller) DoConversion() {
	v1List, err := c.v1Client.List(c.namespace, clients.ListOpts{Ctx: c.ctx})
	if err != nil {
		wrapped := FailedToListGatewayResources(err, "v1", c.namespace)
		contextutils.LoggerFrom(c.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.String("namespace", c.namespace))
	}

	v2alpha1List := make([]*v2alpha1.Gateway, len(v1List), len(v1List))
	for _, oldGateway := range v1List {
		convertedGateway := c.converter.Convert(oldGateway)
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
