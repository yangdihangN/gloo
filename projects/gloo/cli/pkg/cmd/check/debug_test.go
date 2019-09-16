package check_test

import (
	"os"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/check"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
)

var _ = Describe("Debug", func() {
	var (
		vs       *gatewayv1.VirtualService
		vsClient gatewayv1.VirtualServiceClient
	)
	BeforeEach(func() {
		helpers.UseMemoryClients()
		// create a settings object
		vsClient = helpers.MustVirtualServiceClient()
		vs = &gatewayv1.VirtualService{
			Metadata: core.Metadata{
				Name:      "debug",
				Namespace: "gloo-system",
			},
		}
	})

	It("should create a tar file", func() {
		opts := options.Options{}
		opts.Metadata.Namespace = "gloo-system"
		check.DebugResources(&opts)

		_, err := os.Stat(check.Filename)
		Expect(err).NotTo(HaveOccurred())

		os.RemoveAll(check.Filename)
	})

})
