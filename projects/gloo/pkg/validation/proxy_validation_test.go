package validation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	. "github.com/solo-io/gloo/projects/gloo/pkg/validation"
)

var _ = Describe("MakeReport", func() {
	It("generates a report which matches the proxy", func() {
		numListeners := 3
		numVhosts := 5
		numRoutes := 8

		proxy := &v1.Proxy{}
		for i := 0; i < numListeners; i++ {
			httpListener := &v1.HttpListener{}
			proxy.Listeners = append(proxy.Listeners, &v1.Listener{
				ListenerType: &v1.Listener_HttpListener{
					HttpListener: httpListener,
				},
			})

			for j := 0; j < numVhosts; j++ {
				vh := &v1.VirtualHost{}
				httpListener.VirtualHosts = append(httpListener.VirtualHosts, vh)

				for k := 0; k < numRoutes; k++ {
					vh.Routes = append(vh.Routes, &v1.Route{})
				}
			}
		}

		rpt := MakeReport(proxy)
		Expect(rpt.ListenerReports).To(HaveLen(len(proxy.Listeners)))
		for i := range rpt.ListenerReports {
			vhReports := rpt.ListenerReports[i].GetHttpListenerReport().VirtualHostReports
			Expect(vhReports).To(HaveLen(len(proxy.Listeners[i].GetHttpListener().VirtualHosts)))
			for j := range vhReports {
				Expect(vhReports[i].RouteReports).To(HaveLen(len(proxy.Listeners[i].GetHttpListener().VirtualHosts[j].Routes)))

			}
		}
	})
})
