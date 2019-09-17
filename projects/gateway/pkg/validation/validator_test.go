package validation_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2"
	"github.com/solo-io/gloo/projects/gateway/pkg/translator"
	. "github.com/solo-io/gloo/projects/gateway/pkg/validation"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	validationutils "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	"github.com/solo-io/gloo/test/samples"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"google.golang.org/grpc"
)

var _ = Describe("Validator", func() {
	var (
		t  translator.Translator
		vc *mockValidationClient
		ns string
		v  Validator
	)
	BeforeEach(func() {
		t = translator.NewDefaultTranslator()
		vc = &mockValidationClient{}
		ns = "my-namespace"
		v = NewValidator(t, vc, ns)
	})
	It("returns ready == false before sync called", func() {
		Expect(v.Ready()).To(BeFalse())
		Expect(v.ValidateVirtualService(nil, nil)).To(MatchError("Gateway validation is yet not available. Waiting for first snapshot"))
		err := v.Sync(nil, &v2.ApiSnapshot{})
		Expect(err).NotTo(HaveOccurred())
		Expect(v.Ready()).To(BeTrue())
	})

	Context("validating a route table", func() {
		Context("proxy validation accepted", func() {
			It("accepts the rt", func() {
				vc.validateProxy = acceptProxy
				us := samples.SimpleUpstream()
				snap := samples.GatewaySnapshotWithDelegates(us.Metadata.Ref(), ns)
				err := v.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())
				err = v.ValidateRouteTable(context.TODO(), snap.RouteTables[0])
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("rendered proxy had errors"))
			})
		})
	})

	FIt("write me a boy", func() {

		cfg, err := kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())

		cache := kube.NewKubeCache(context.TODO())
		rtc, _ := gatewayv1.NewRouteTableClient(&factory.KubeResourceClientFactory{
			Crd:         gatewayv1.RouteTableCrd,
			Cfg:         cfg,
			SharedCache: cache,
		})
		snap := samples.GatewaySnapshotWithMultiDelegates(core.ResourceRef{"a", "b"}, "gloo-system")
		snap.RouteTables.Each(func(element *gatewayv1.RouteTable) {
			_, err = rtc.Write(element, clients.WriteOpts{})
		})

		vsc := helpers.MustVirtualServiceClient()
		snap.VirtualServices.Each(func(element *gatewayv1.VirtualService) {
			_, err = vsc.Write(element, clients.WriteOpts{})
		})
		//Expect(err).NotTo(HaveOccurred())

		cfg.Timeout = 0
	})

	Context("validating a virtual service", func() {

		Context("proxy validation returns error", func() {
			It("rejects the vs", func() {
				vc.validateProxy = failProxy
				us := samples.SimpleUpstream()
				snap := samples.SimpleGatewaySnapshot(us.Metadata.Ref(), ns)
				err := v.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())
				err = v.ValidateVirtualService(context.TODO(), snap.VirtualServices[0])
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("rendered proxy had errors"))
			})
		})
		Context("proxy validation accepted", func() {
			It("accepts the vs", func() {
				vc.validateProxy = acceptProxy
				us := samples.SimpleUpstream()
				snap := samples.SimpleGatewaySnapshot(us.Metadata.Ref(), ns)
				err := v.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())
				err = v.ValidateVirtualService(context.TODO(), snap.VirtualServices[0])
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("no gateways for virtualservice", func() {
			It("accepts the vs", func() {
				vc.validateProxy = failProxy
				us := samples.SimpleUpstream()
				snap := samples.SimpleGatewaySnapshot(us.Metadata.Ref(), ns)
				snap.Gateways.Each(func(element *v2.Gateway) {
					http, ok := element.GatewayType.(*v2.Gateway_HttpGateway)
					if !ok {
						return
					}
					http.HttpGateway.VirtualServiceSelector = map[string]string{"nobody": "hastheselabels"}

				})
				err := v.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())
				err = v.ValidateVirtualService(context.TODO(), snap.VirtualServices[0])
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("virtual service rejected", func() {
			It("rejects the vs", func() {
				badRoute := &gatewayv1.Route{
					Action: &gatewayv1.Route_DelegateAction{
						DelegateAction: nil,
					},
				}

				// validate proxy should never be called
				vc.validateProxy = nil
				us := samples.SimpleUpstream()
				snap := samples.SimpleGatewaySnapshot(us.Metadata.Ref(), ns)
				snap.VirtualServices[0].VirtualHost.Routes = append(snap.VirtualServices[0].VirtualHost.Routes, badRoute)
				err := v.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())
				err = v.ValidateVirtualService(context.TODO(), snap.VirtualServices[0])
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("could not render proxy from *v1.VirtualService"))
			})
		})
	})
})

type mockValidationClient struct {
	validateProxy func(ctx context.Context, in *v1.Proxy, opts ...grpc.CallOption) (*validation.ProxyReport, error)
}

func (c *mockValidationClient) ValidateProxy(ctx context.Context, in *v1.Proxy, opts ...grpc.CallOption) (*validation.ProxyReport, error) {
	return c.validateProxy(ctx, in, opts...)
}

func acceptProxy(ctx context.Context, in *v1.Proxy, opts ...grpc.CallOption) (report *validation.ProxyReport, e error) {
	return validationutils.MakeReport(in), nil
}

func failProxy(ctx context.Context, in *v1.Proxy, opts ...grpc.CallOption) (report *validation.ProxyReport, e error) {
	rpt := validationutils.MakeReport(in)
	validationutils.AppendListenerError(rpt.ListenerReports[0], validation.ListenerReport_Error_SSLConfigError, "")
	return rpt, nil
}
