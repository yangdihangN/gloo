package validation_test

import (
	"context"

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
	Context("validates a virtual service", func() {
		t := translator.NewDefaultTranslator()
		vc := &mockValidationClient{}
		ns := "my-namespace"
		v := NewValidator(t, vc, ns)

		It("returns ready == false before sync called", func() {
			Expect(v.Ready()).To(BeFalse())
			v.Sync(nil, &v2.ApiSnapshot{})
			Expect(v.Ready()).To(BeTrue())
		})
	})
})

type mockValidationClient struct {
	validateProxy func(ctx context.Context, in *v1.Proxy, opts ...grpc.CallOption) (*validation.ProxyReport, error)
}

func (c *mockValidationClient) ValidateProxy(ctx context.Context, in *v1.Proxy, opts ...grpc.CallOption) (*validation.ProxyReport, error) {
	return c.validateProxy(ctx, in, opts...)
}
