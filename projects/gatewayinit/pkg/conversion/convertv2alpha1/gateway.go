package convertv2alpha1

import (
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
)

type GatewayConverter interface {
	Convert(existing *v1.Gateway) *v2alpha1.Gateway
}

type gatewayConverter struct {
}

func NewGatewayConverter() GatewayConverter {
	return &gatewayConverter{}
}

func (c *gatewayConverter) Convert(existing *v1.Gateway) *v2alpha1.Gateway {
	return &v2alpha1.Gateway{
		Ssl:           existing.Ssl,
		BindAddress:   existing.BindAddress,
		BindPort:      existing.BindPort,
		UseProxyProto: existing.UseProxyProto,
		GatewayType: &v2alpha1.Gateway_HttpGateway{
			HttpGateway: &v2alpha1.HttpGateway{
				VirtualServices: existing.VirtualServices,
				Plugins:         existing.Plugins,
			},
		},
	}
}
