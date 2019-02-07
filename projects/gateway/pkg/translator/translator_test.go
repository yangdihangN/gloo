package translator_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/solo-io/gloo/projects/gateway/pkg/translator"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const ns = "gloo-system"

var _ = Describe("Translator", func() {
	var (
		snap *v1.ApiSnapshot
	)
	BeforeEach(func() {
		snap = &v1.ApiSnapshot{
			Gateways: v1.GatewaysByNamespace{
				ns: v1.GatewayList{{Metadata: core.Metadata{Namespace: ns, Name: "name"}}},
			},
			VirtualServices: v1.VirtualServicesByNamespace{
				ns: v1.VirtualServiceList{{Metadata: core.Metadata{Namespace: ns, Name: "name"}}, {Metadata: core.Metadata{Namespace: ns, Name: "name2"}}},
			},
		}
	})

	It("should translate proxy with default name", func() {

		proxylist, _ := Translate(context.Background(), ns, snap)

		Expect(proxylist).To(HaveLen(1))
		proxy := proxylist[0].Proxy
		Expect(proxy.Metadata.Name).To(Equal(DefaultProxyName))
		Expect(proxy.Metadata.Namespace).To(Equal(ns))
	})

	It("should translate an empty gateway to have all vservices", func() {

		proxylist, errs := Translate(context.Background(), ns, snap)

		Expect(proxylist).To(HaveLen(1))
		Expect(errs).To(HaveLen(3))
		proxy := proxylist[0].Proxy
		Expect(proxy.Listeners).To(HaveLen(1))
		listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
		Expect(listener.VirtualHosts).To(HaveLen(2))
	})

	It("should translate an gateway to only have its vservices", func() {
		snap.Gateways[ns][0].VirtualServices = []core.ResourceRef{snap.VirtualServices[ns][0].Metadata.Ref()}

		proxylist, _ := Translate(context.Background(), ns, snap)

		Expect(proxylist).To(HaveLen(1))
		proxy := proxylist[0].Proxy
		Expect(proxy.Listeners).To(HaveLen(1))
		listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
		Expect(listener.VirtualHosts).To(HaveLen(1))
	})

	It("should translate two gateways with different proxy names to two proxies", func() {
		snap.Gateways[ns] = append(snap.Gateways[ns], &v1.Gateway{Metadata: core.Metadata{Namespace: ns, Name: "name"}, ProxyName: "1"})

		proxylist, _ := Translate(context.Background(), ns, snap)

		Expect(proxylist).To(HaveLen(2))
		Expect(proxylist[0].Proxy.Metadata.Name).To(Equal(DefaultProxyName))
		Expect(proxylist[0].Proxy.Metadata.Namespace).To(Equal(ns))

		Expect(proxylist[1].Proxy.Metadata.Name).To(Equal("1"))
		Expect(proxylist[1].Proxy.Metadata.Namespace).To(Equal(ns))
	})

})
