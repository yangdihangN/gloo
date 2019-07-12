package conversion

import (
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type GatewayUpConverter struct{}

func (u GatewayUpConverter) FromV1ToV2Alpha1(src *v1.Gateway) *v2alpha1.Gateway {
	return &v2alpha1.Gateway{
		Metadata: core.Metadata{
			Namespace: src.GetMetadata().Namespace,
			Name:      src.GetMetadata().Name,
		},
		Ssl:           src.Ssl,
		BindAddress:   src.BindAddress,
		BindPort:      src.BindPort,
		UseProxyProto: src.UseProxyProto,
		GatewayType: &v2alpha1.Gateway_HttpGateway{
			HttpGateway: &v2alpha1.HttpGateway{
				VirtualServices: src.VirtualServices,
				Plugins:         src.Plugins,
			},
		},
	}
}

type GatewayDownConverter struct{}

func (u GatewayUpConverter) FromV2Alpha1ToV1(src *v2alpha1.Gateway) *v1.Gateway {
	return &v1.Gateway{
		Metadata: core.Metadata{
			Namespace: src.GetMetadata().Namespace,
			Name:      src.GetMetadata().Name,
		},
		Ssl:             src.Ssl,
		BindAddress:     src.BindAddress,
		BindPort:        src.BindPort,
		UseProxyProto:   src.UseProxyProto,
		VirtualServices: src.GetHttpGateway().GetVirtualServices(),
		Plugins:         src.GetHttpGateway().GetPlugins(),
	}
}
