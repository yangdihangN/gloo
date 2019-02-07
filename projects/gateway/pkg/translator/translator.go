package translator

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/solo-io/solo-kit/pkg/utils/contextutils"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	glooutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
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
			mergedVirtualServices := merge(gateway, virtualServices, proxyResourceErrs)

			listener := desiredListener(gateway, mergedVirtualServices)
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

func validateGateways(gateways v1.GatewayList, resourceErrs reporter.ResourceErrors) {
	bindAddresses := map[string]v1.GatewayList{}
	// if two gateway (=listener) that belong to the same proxy share the same bind address,
	// they are invalid.
	for _, gw := range gateways {
		bindAddress := fmt.Sprintf("%s:%d", gw.BindAddress, gw.BindPort)
		bindAddresses[bindAddress] = append(bindAddresses[bindAddress], gw)
	}

	for addr, gateways := range bindAddresses {
		if len(gateways) > 1 {
			for _, gw := range gateways {
				resourceErrs.AddError(gw, fmt.Errorf("bind-addres %s is not unique in a proxy. gateways: %s", addr, strings.Join(gatewaysRefsToString(gateways), ",")))
			}
		}
	}
}

func gatewaysRefsToString(gateways v1.GatewayList) []string {
	var ret []string
	for _, gw := range gateways {
		ret = append(ret, gw.Metadata.Ref().Key())
	}
	return ret
}

func merge(gateway *v1.Gateway, virtualServices v1.VirtualServiceList, resourceErrs reporter.ResourceErrors) v1.VirtualServiceList {

	return nil
}

func validateVirtualServices(gateway *v1.Gateway, virtualServices v1.VirtualServiceList, resourceErrs reporter.ResourceErrors) {

	domainKeysSets := map[string]v1.VirtualServiceList{}
	for _, vs := range virtualServices {
		if vs.VirtualHost == nil {
			continue
		}
		domainsKey := domainsToKey(vs.VirtualHost.Domains)
		domainKeysSets[domainsKey] = append(domainKeysSets[domainsKey], vs)
	}

	domainSet := map[string][]string{}
	// make sure each domain is only in one domain set
	for k, vslist := range domainKeysSets {
		// take the first one as they are all the same
		domains := vslist[0].VirtualHost.Domains
		for _, d := range domains {
			domainSet[d] = append(domainSet[d], k)
		}
	}

	// report errors
	for domain, domainSetKeys := range domainSet {
		if len(domainSetKeys) > 1 {
			resourceErrs.AddError(gateway, fmt.Errorf("domain %s is present in more than one vservice set in this gateway", domain))
		}
	}
	// return merged list
	var ret v1.VirtualServiceList
	for _, vslist := range domainKeysSets {
		// take the first one as they are all the same
		var routes []*gloov1.Route
		for _, vs := range vslist {
			routes = append(routes, vs.VirtualHost.Routes...)
		}
		glooutils.SortRoutesByPath(routes)
		mergedVs := &v1.VirtualService{
			VirtualHost: &gloov1.VirtualHost{
				Domains:            vslist[0].VirtualHost.Domains,
				Routes:             routes,
				Name:               "TODO get merged name",
				VirtualHostPlugins: nil, /* MAJOR TODO what todo!?*/
			},
			SslConfig: nil,             /* MAJOR TODO what todo!?*/
			Metadata:  core.Metadata{}, /* MAJOR TODO what todo!?*/
		}
		ret = append(ret, mergedVs)
	}

}

func domainsToKey(domains []string) string {
	// copy before mutating for good measure
	domains = append([]string{}, domains...)
	// sort, and join all domains with an out of band character, like |
	sort.Strings(domains)
	return strings.Join(domains, "|")
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
	// TODO: fix with ssl
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
