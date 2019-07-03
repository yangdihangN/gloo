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

func (c *controller) DoConversion() {
	v1List, err := c.v1Client.List(c.namespace, clients.ListOpts{Ctx: c.ctx})
	if err != nil {
		wrapped := FailedToListGatewayResources(err, "v1", c.namespace)
		contextutils.LoggerFrom(c.ctx).Errorw(wrapped.Error(), zap.Error(err))
	}

	v2alpha1List := make([]*v2alpha1.Gateway, len(v1List), len(v1List))
	for _, oldGateway := range v1List {
		convertedGateway, err := c.converter.Convert(oldGateway)
		if err != nil {
			// TODO Log ?
		} else {
			if err := c.v1Client.Delete(
				convertedGateway.GetMetadata().GetNamespace(),
				convertedGateway.GetMetadata().GetName(),
				clients.DeleteOpts{Ctx: c.ctx}); err != nil {

				// TODO LOG
			}

			v2alpha1List = append(v2alpha1List, convertedGateway)
		}
	}

	for _, newGateway := range v2alpha1List {
		_, err := c.v2alpha1Client.Write(newGateway, clients.WriteOpts{Ctx: c.ctx})
		if err != nil {
			// TODO LOG
		}
	}
}
