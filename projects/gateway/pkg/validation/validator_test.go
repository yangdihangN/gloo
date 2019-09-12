package validation_test

import (
	"context"
	validationutils "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	"github.com/solo-io/gloo/test/samples"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2"
	"github.com/solo-io/gloo/projects/gateway/pkg/translator"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"google.golang.org/grpc"

	. "github.com/solo-io/gloo/projects/gateway/pkg/validation"
)

var _ = Describe("Validator", func() {
	Context("validating a virtual service", func() {
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

		Context("proxy validation returns error", func() {
			It("rejects the vs", func() {
				vc.validateProxy = failProxy
				us := samples.SimpleUpstream()
				snap := samples.SimpleGatewaySnapshot(us.Metadata.Ref(), ns)
				err := v.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())
				err = v.ValidateVirtualService(context.TODO(), snap.VirtualServices[0])
				Expect(err).To(HaveOccurred())
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
		Context("proxy rejected", func() {
			It("rejects the vs", func() {
				vc.validateProxy = failProxy
				us := samples.SimpleUpstream()
				snap := samples.SimpleGatewaySnapshot(us.Metadata.Ref(), ns)
				err := v.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())
				err = v.ValidateVirtualService(context.TODO(), snap.VirtualServices[0])
				Expect(err).To(HaveOccurred())
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
