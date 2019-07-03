package conversion

import (
	"context"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

type GatewayConverter interface {
	Convert(existing *v1.Gateway) *v2alpha1.Gateway
}

type gatewayConverter struct {
	ctx context.Context
}

func (c *gatewayConverter) Convert(existing *v1.Gateway) *v2alpha1.Gateway {
	return &v2alpha1.Gateway{
		Ssl:         existing.Ssl,
		BindAddress: existing.BindAddress,
		BindPort:    existing.BindPort,
		Plugins:     &gloov1.ListenerPlugins{},
	}
}
