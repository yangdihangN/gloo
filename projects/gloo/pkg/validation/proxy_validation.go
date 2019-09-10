package validation

import (
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

func MakeReport(proxy *v1.Proxy) *validation.ProxyReport {
	listeners := proxy.GetListeners()
	listenerReports := make([]*validation.ListenerReport, len(listeners))

	for i, listener := range listeners {
		switch listenerType := listener.GetListenerType().(type) {
		case *v1.Listener_HttpListener:

			vhostReports := make([]*validation.VirtualHostReport, len(listenerType.HttpListener.GetVirtualHosts()))

			for j, vh := range listenerType.HttpListener.GetVirtualHosts() {
				routeReports := make([]*validation.RouteReport, len(vh.GetRoutes()))

				vhostReports[j] = &validation.VirtualHostReport{
					RouteReports: routeReports,
				}
			}

			listenerReports[i] = &validation.ListenerReport{
				ListenerTypeReport: &validation.ListenerReport_HttpListenerReport{
					HttpListenerReport: &validation.HttpListenerReport{
						VirtualHostReports: vhostReports,
					},
				},
			}
		case *v1.Listener_TcpListener:
			tcpHostReports := make([]*validation.TcpHostReport, len(listenerType.TcpListener.GetTcpHosts()))
			listenerReports[i] = &validation.ListenerReport{
				ListenerTypeReport: &validation.ListenerReport_TcpListenerReport{
					TcpListenerReport: &validation.TcpListenerReport{
						TcpHostReports: tcpHostReports,
					},
				},
			}
		}
	}

	return &validation.ProxyReport{
		ListenerReports: listenerReports,
	}
}
