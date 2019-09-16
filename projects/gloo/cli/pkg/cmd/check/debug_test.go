package check_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/check"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
)

var _ = Describe("Debug", func() {
	BeforeEach(func() {
		helpers.UseMemoryClients()
	})

	It("should create a tar file", func() {
		opts := options.Options{}
		opts.Metadata.Namespace = "gloo-system"
		opts.Top.File = "/tmp/log.tgz"
		check.DebugResources(&opts)

		_, err := os.Stat(opts.Top.File)
		Expect(err).NotTo(HaveOccurred())

		os.RemoveAll(opts.Top.File)
	})

})
