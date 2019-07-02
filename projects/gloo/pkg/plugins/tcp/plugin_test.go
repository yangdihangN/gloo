package tcp_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/tcp"

	. "github.com/solo-io/gloo/projects/gloo/pkg/plugins/tcp"
	translatorutil "github.com/solo-io/gloo/projects/gloo/pkg/translator"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoylistener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	envoytcp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	envoyutil "github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

var _ = Describe("Plugin", func() {
	var (
		in      *v1.Listener
		outl    *envoyapi.Listener
		filters []envoylistener.Filter
		tcps    *tcp.TcpProxySettings
	)

	BeforeEach(func() {
		pd := func(t time.Duration) *time.Duration { return &t }
		tcps = &tcp.TcpProxySettings{
			MaxConnectAttempts: &types.UInt32Value{
				Value: 5,
			},
			IdleTimeout: pd(5 * time.Second),
		}
		tl := &v1.TcpListener{
			ListenerPlugins: &v1.TcpListenerPlugins{
				TcpProxySettings: tcps,
			},
		}

		in = &v1.Listener{
			ListenerType: &v1.Listener_TcpListener{
				TcpListener: tl,
			},
		}
		filters = []envoylistener.Filter{{
			Name: envoyutil.TCPProxy,
		}}
		outl = &envoyapi.Listener{
			FilterChains: []envoylistener.FilterChain{{
				Filters: filters,
			}},
		}
	})
	It("copy all settings to tcp filter", func() {

		p := NewPlugin()
		err := p.ProcessListener(plugins.Params{}, in, outl)
		Expect(err).NotTo(HaveOccurred())

		var cfg envoytcp.TcpProxy
		err = translatorutil.ParseConfig(&filters[0], &cfg)
		Expect(err).NotTo(HaveOccurred())

		Expect(cfg.IdleTimeout).To(Equal(tcps.IdleTimeout))
		Expect(cfg.MaxConnectAttempts).To(Equal(tcps.MaxConnectAttempts))
	})

	It("appends the tls inspector listener filter", func() {
		p := NewPlugin()
		err := p.ProcessListener(plugins.Params{}, in, outl)
		Expect(err).NotTo(HaveOccurred())
		Expect(outl.ListenerFilters).To(HaveLen(1))
	})

	It("does not append the tls inspector if it already exists", func() {
		p := NewPlugin()
		outl.ListenerFilters = append(outl.ListenerFilters, envoylistener.ListenerFilter{
			Name: envoyutil.TlsInspector,
		})
		err := p.ProcessListener(plugins.Params{}, in, outl)
		Expect(err).NotTo(HaveOccurred())
		Expect(outl.ListenerFilters).To(HaveLen(1))
	})

})
