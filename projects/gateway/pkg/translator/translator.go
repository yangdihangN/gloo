package translator

import (
	"context"
	"fmt"

	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const GatewayProxyName = "gateway-proxy"

type Translator interface {
	Translate(ctx context.Context, namespace string, snap *v2alpha1.ApiSnapshot) (*gloov1.Proxy, reporter.ResourceErrors)
}

type translator struct {
	factories []ListenerFactory
}

func NewTranslator(factories []ListenerFactory) *translator {
	return &translator{factories: factories}
}

type ListenerFactory interface {
	GenerateListeners(snap *v2alpha1.ApiSnapshot, filteredGateways []*v2alpha1.Gateway, resourceErrs reporter.ResourceErrors) []*gloov1.Listener
}

func (t *translator) Translate(ctx context.Context, namespace string, snap *v2alpha1.ApiSnapshot) (*gloov1.Proxy, reporter.ResourceErrors) {
	logger := contextutils.LoggerFrom(ctx)

	filteredGateways := filterGatewaysForNamespace(snap.Gateways, namespace)

	resourceErrs := make(reporter.ResourceErrors)
	resourceErrs.Accept(filteredGateways.AsInputResources()...)
	resourceErrs.Accept(snap.VirtualServices.AsInputResources()...)
	if len(filteredGateways) == 0 {
		logger.Debugf("%v had no gateways", snap.Hash())
		return nil, resourceErrs
	}
	if len(snap.VirtualServices) == 0 {
		logger.Debugf("%v had no virtual services", snap.Hash())
		return nil, resourceErrs
	}
	validateGateways(filteredGateways, resourceErrs)
	var listeners []*gloov1.Listener
	for _, factory := range t.factories {
		listeners = append(listeners, factory.GenerateListeners(snap, filteredGateways, resourceErrs)...)
	}
	if len(listeners) != len(filteredGateways) {
		logger.Debug("length of listeners does not match the input gateways")
	}
	return &gloov1.Proxy{
		Metadata: core.Metadata{
			Name:      GatewayProxyName,
			Namespace: namespace,
		},
		Listeners: listeners,
	}, resourceErrs
}

func standardListener(gateway *v2alpha1.Gateway) *gloov1.Listener {
	return &gloov1.Listener{
		Name:            gatewayName(gateway),
		BindAddress:     gateway.BindAddress,
		BindPort:        gateway.BindPort,
		ListenerPlugins: gateway.Plugins,
		UseProxyProto:   gateway.UseProxyProto,
	}
}

func gatewayName(gateway *v2alpha1.Gateway) string {
	return fmt.Sprintf("listener-%s-%d", gateway.BindAddress, gateway.BindPort)
}



type TcpTranslator struct {}

func (t *TcpTranslator) GenerateListeners(snap *v2alpha1.ApiSnapshot, filteredGateways []*v2alpha1.Gateway, resourceErrs reporter.ResourceErrors) []*gloov1.Listener {
	var result []*gloov1.Listener
	for _, gateway := range filteredGateways {
		tcpGateway := gateway.GetTcpGateway()
		if tcpGateway == nil {
			continue
		}
		listener := standardListener(gateway)
		listener.ListenerType = &gloov1.Listener_TcpListener{
			TcpListener: &gloov1.TcpListener{
				ListenerPlugins: tcpGateway.Plugins,
			},
		}
		result = append(result, listener)
	}
	return result
}

