package convertgateway

import (
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
)

type V2alpha1Converter interface {
	Convert(existing *v1.Gateway) *v2alpha1.Gateway
}

type v2alpha1Converter struct{}

func NewV2alpha1Converter() V2alpha1Converter {
	return &v2alpha1Converter{}
}

func (c *v2alpha1Converter) Convert(existing *v1.Gateway) *v2alpha1.Gateway {
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
