package translator

import (
	"sort"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoyauth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoylistener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	types "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/go-utils/contextutils"
)

func (t *translator) computeListener(params plugins.Params, proxy *v1.Proxy, listener *v1.Listener, translatorReport reportFunc) *envoyapi.Listener {
	params.Ctx = contextutils.WithLogger(params.Ctx, "compute_listener."+listener.Name)

	report := func(err error, format string, args ...interface{}) {
		translatorReport(err, "listener."+format, args...)
	}
	validateListenerPorts(proxy, report)
	var filterChains []envoylistener.FilterChain
	switch listenerType := listener.GetListenerType().(type) {
	case *v1.Listener_HttpListener:
		listenerFilters := t.computeListenerFilters(params, listener, report)
		if len(listenerFilters) == 0 {
			return nil
		}
		t.httpConnectionManager(params, listener, listenerFilters, report)
		computeFilterChainsFromSslConfig(params.Snapshot, listener, sortListenerFilters(listenerFilters), report)
	case *v1.Listener_TcpListener:
		// run the Listener Plugins
		for _, plug := range t.plugins {
			listenerPlugin, ok := plug.(plugins.ListenerFilterChainPlugin)
			if !ok {
				continue
			}
			result, err := listenerPlugin.ProcessListenerFilterChain(params, listener)
			if err != nil {
				report(err, "plugin error on listener filter chain")
				continue
			}
			filterChains = append(filterChains, result...)
		}
		filterChains = t.tcpFilterChains(params, listener, listenerType.TcpListener, report)
	}

	out := &envoyapi.Listener{
		Name: listener.Name,
		Address: envoycore.Address{
			Address: &envoycore.Address_SocketAddress{
				SocketAddress: &envoycore.SocketAddress{
					Protocol: envoycore.TCP,
					Address:  listener.BindAddress,
					PortSpecifier: &envoycore.SocketAddress_PortValue{
						PortValue: listener.BindPort,
					},
					Ipv4Compat: true,
				},
			},
		},
		FilterChains: filterChains,
	}

	// run the Listener Plugins
	for _, plug := range t.plugins {
		listenerPlugin, ok := plug.(plugins.ListenerPlugin)
		if !ok {
			continue
		}
		if err := listenerPlugin.ProcessListener(params, listener, out); err != nil {
			report(err, "plugin error on listener")
		}
	}

	return out
}

func (t *translator) tcpFilterChains(params plugins.Params, listener *v1.Listener, tcpListener *v1.TcpListener, report reportFunc) []envoylistener.FilterChain {
	var filterChains []envoylistener.FilterChain
	for _, tcpHost := range tcpListener.TcpHosts {
		listenerFilters := t.computeListenerFilters(params, listener, report)
		if len(listenerFilters) == 0 {
			// nothing to do, return nil
			return nil
		}

		filterChain, err := computerTcpFilterChain(params.Snapshot, listener, sortListenerFilters(listenerFilters), tcpHost)
		if err != nil {
			report(err, "could not compute tcp filter chain for %v", tcpHost)
			continue
		}
		filterChains = append(filterChains, filterChain)
	}
	return filterChains
}

func (t *translator) computeListenerFilters(params plugins.Params, listener *v1.Listener, report reportFunc) []plugins.StagedListenerFilter {
	var listenerFilters []plugins.StagedListenerFilter
	// run the Listener Filter Plugins
	for _, plug := range t.plugins {
		filterPlugin, ok := plug.(plugins.ListenerFilterPlugin)
		if !ok {
			continue
		}
		stagedFilters, err := filterPlugin.ProcessListenerFilter(params, listener)
		if err != nil {
			report(err, "listener plugin error")
		}
		for _, listenerFilter := range stagedFilters {
			listenerFilters = append(listenerFilters, listenerFilter)
		}
	}
	return listenerFilters
}

func (t *translator) httpConnectionManager(params plugins.Params, listener *v1.Listener, listenerFilters []plugins.StagedListenerFilter, report reportFunc) {
	// add the http connection manager if listener is HTTP and has >= 1 virtual hosts
	httpListener, ok := listener.ListenerType.(*v1.Listener_HttpListener)
	if !ok || len(httpListener.HttpListener.VirtualHosts) == 0 {
		return
	}

	// add the http connection manager filter after all the InAuth Listener Filters
	rdsName := routeConfigName(listener)
	httpConnMgr := t.computeHttpConnectionManagerFilter(params, httpListener.HttpListener, rdsName, report)
	listenerFilters = append(listenerFilters, plugins.StagedListenerFilter{
		ListenerFilter: httpConnMgr,
		Stage:          plugins.PostInAuth,
	})
}

// create a duplicate of the listener filter chain for each ssl cert we want to serve
// if there is no SSL config on the listener, the envoy listener will have one insecure filter chain
func computeFilterChainsFromSslConfig(snap *v1.ApiSnapshot, listener *v1.Listener, listenerFilters []envoylistener.Filter, report reportFunc) []envoylistener.FilterChain {

	// if no ssl config is provided, return a single insecure filter chain
	if len(listener.SslConfigurations) == 0 {
		return []envoylistener.FilterChain{{
			Filters:       listenerFilters,
			UseProxyProto: listener.UseProxyProto,
		}}
	}

	var secureFilterChains []envoylistener.FilterChain

	sslCfgTranslator := utils.NewSslConfigTranslator(snap.Secrets)
	for _, sslConfig := range listener.SslConfigurations {
		// get secrets
		downstreamConfig, err := sslCfgTranslator.ResolveDownstreamSslConfig(sslConfig)
		if err != nil {
			report(err, "invalid secrets for listener %v", listener.Name)
			continue
		}
		filterChain := newSslFilterChain(downstreamConfig, sslConfig.SniDomains, listener.UseProxyProto, listenerFilters)
		secureFilterChains = append(secureFilterChains, filterChain)
	}
	return secureFilterChains
}

// create a duplicate of the listener filter chain for each ssl cert we want to serve
// if there is no SSL config on the listener, the envoy listener will have one insecure filter chain
func computerTcpFilterChain(snap *v1.ApiSnapshot, listener *v1.Listener, listenerFilters []envoylistener.Filter, host *v1.TcpHost) (envoylistener.FilterChain, error) {
	sslConfig := host.GetSslConfig()
	if sslConfig == nil {
		return envoylistener.FilterChain{
			Filters:       listenerFilters,
			UseProxyProto: listener.UseProxyProto,
		}, nil
	}

	sslCfgTranslator := utils.NewSslConfigTranslator(snap.Secrets)
	downstreamConfig, err := sslCfgTranslator.ResolveDownstreamSslConfig(sslConfig)
	if err != nil {
		return envoylistener.FilterChain{}, errors.Wrapf(err, "invalid secrets for listener %v", listener.Name)
	}
	return newSslFilterChain(downstreamConfig, sslConfig.SniDomains, listener.UseProxyProto, listenerFilters), nil
}

func validateListenerPorts(proxy *v1.Proxy, report reportFunc) {
	listenersByPort := make(map[uint32][]string)
	for _, listener := range proxy.Listeners {
		listenersByPort[listener.BindPort] = append(listenersByPort[listener.BindPort], listener.Name)
	}
	for port, listeners := range listenersByPort {
		if len(listeners) == 1 {
			continue
		}
		report(errors.Errorf("port %v is shared by listeners %v", port, listeners), "invalid listener config")
	}
}

func newSslFilterChain(downstreamConfig *envoyauth.DownstreamTlsContext, sniDomains []string, useProxyProto *types.BoolValue, listenerFilters []envoylistener.Filter) envoylistener.FilterChain {

	return envoylistener.FilterChain{
		FilterChainMatch: &envoylistener.FilterChainMatch{
			ServerNames: sniDomains,
		},
		Filters:       listenerFilters,
		TlsContext:    downstreamConfig,
		UseProxyProto: useProxyProto,
	}
}

func sortListenerFilters(filters []plugins.StagedListenerFilter) []envoylistener.Filter {
	// sort them first by stage, then by name.
	less := func(i, j int) bool {
		filteri := filters[i]
		filterj := filters[j]
		if filteri.Stage != filterj.Stage {
			return filteri.Stage < filterj.Stage
		}
		return filteri.ListenerFilter.Name < filterj.ListenerFilter.Name
	}
	sort.SliceStable(filters, less)

	var sortedFilters []envoylistener.Filter
	for _, filter := range filters {
		sortedFilters = append(sortedFilters, filter.ListenerFilter)
	}

	return sortedFilters
}
