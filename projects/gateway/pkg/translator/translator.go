package translator

import (
	"context"
	"fmt"

	"github.com/solo-io/solo-kit/pkg/utils/contextutils"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const DefaultProxyName = "gateway-proxy"

type ProxyWithResourceErrors struct {
	Proxy          *gloov1.Proxy
	ResourceErrors reporter.ResourceErrors
}

func Translate(ctx context.Context, namespace string, snap *v1.ApiSnapshot) ([]ProxyWithResourceErrors, reporter.ResourceErrors) {
	logger := contextutils.LoggerFrom(ctx)

	resourceErrs := make(reporter.ResourceErrors)
	resourceErrs.Accept(snap.Gateways.List().AsInputResources()...)
	resourceErrs.Accept(snap.VirtualServices.List().AsInputResources()...)
	if len(snap.Gateways.List()) == 0 {
		logger.Debugf("%v had no gateways", snap.Hash())
		return nil, resourceErrs
	}
	if len(snap.VirtualServices.List()) == 0 {
		logger.Debugf("%v had no virtual services", snap.Hash())
		return nil, resourceErrs
	}

	var proxiesAndErrors []ProxyWithResourceErrors

	for proxyName, gateways := range groupGatwaysPerProxy(snap.Gateways.List()) {
		proxyResourceErrs := make(reporter.ResourceErrors)
		proxyResourceErrs.Accept(gateways.AsInputResources()...)

		validateGateways(gateways, proxyResourceErrs)
		var listeners []*gloov1.Listener
		for _, gateway := range gateways {
			virtualServices := getVirtualServiceForGateway(gateway, snap.VirtualServices.List(), proxyResourceErrs)
			proxyResourceErrs.Accept(virtualServices.AsInputResources()...)
			validateVirtualServices(gateway, virtualServices, proxyResourceErrs)

			listener := desiredListener(gateway, virtualServices)
			listeners = append(listeners, listener)
		}
		proxy := &gloov1.Proxy{
			Metadata: core.Metadata{
				Name:      proxyName,
				Namespace: namespace,
			},
			Listeners: listeners,
		}

		for k, v := range proxyResourceErrs {
			resourceErrs.AddError(k, v)
		}

		proxiesAndErrors = append(proxiesAndErrors, ProxyWithResourceErrors{Proxy: proxy, ResourceErrors: proxyResourceErrs})
	}

	return proxiesAndErrors, resourceErrs
}

// TODO(ilackarms): implement validation func
func validateGateways(gateways v1.GatewayList, resourceErrs reporter.ResourceErrors) {

}

func validateVirtualServices(gateways *v1.Gateway, virtualServices v1.VirtualServiceList, resourceErrs reporter.ResourceErrors) {

}

func groupGatwaysPerProxy(gatewayList v1.GatewayList) map[string]v1.GatewayList {
	proxyToGateway := make(map[string]v1.GatewayList)

	for _, gw := range gatewayList {
		name := gw.ProxyName
		if name == "" {
			name = DefaultProxyName
		}
		proxyToGateway[name] = append(proxyToGateway[name], gw)
	}

	return proxyToGateway
}

func getVirtualServiceForGateway(gateway *v1.Gateway, virtualServices v1.VirtualServiceList, resourceErrs reporter.ResourceErrors) v1.VirtualServiceList {
	virtualServicesForGateway := gateway.VirtualServices
	// add all virtual services if empty
	if len(gateway.VirtualServices) == 0 {
		for _, virtualService := range virtualServices {
			virtualServicesForGateway = append(virtualServicesForGateway, core.ResourceRef{
				Name:      virtualService.GetMetadata().Name,
				Namespace: virtualService.GetMetadata().Namespace,
			})
		}
	}

	var ret v1.VirtualServiceList
	for _, ref := range virtualServicesForGateway {
		// virtual service must live in the same namespace as gateway
		virtualService, err := virtualServices.Find(ref.Strings())
		if err != nil {
			resourceErrs.AddError(gateway, err)
			continue
		}
		ret = append(ret, virtualService)
	}
	return ret
}

func desiredListener(gateway *v1.Gateway, virtualServicesForGateway v1.VirtualServiceList) *gloov1.Listener {

	var (
		virtualHosts []*gloov1.VirtualHost
		sslConfigs   []*gloov1.SslConfig
	)

	for _, virtualService := range virtualServicesForGateway {
		ref := virtualService.Metadata.Ref()
		if virtualService.VirtualHost == nil {
			virtualService.VirtualHost = &gloov1.VirtualHost{}
		}
		virtualService.VirtualHost.Name = fmt.Sprintf("%v.%v", ref.Namespace, ref.Name)
		virtualHosts = append(virtualHosts, virtualService.VirtualHost)
	}
	return &gloov1.Listener{
		Name:        gateway.Metadata.Name,
		BindAddress: gateway.BindAddress,
		BindPort:    gateway.BindPort,
		ListenerType: &gloov1.Listener_HttpListener{
			HttpListener: &gloov1.HttpListener{
				VirtualHosts: virtualHosts,
			},
		},
		SslConfiguations: sslConfigs,
	}
}
