package cli_unit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install"
)

var _ = FDescribe("Manifest", func() {

	var (
		initial = []byte("spec:\ncontainers:\n- image: soloio/gloo:0.5.5-2-g60e818f7\nimagePullPolicy: Always\nname: gloo\nports:")
		updated = []byte("spec:\ncontainers:\n- image: soloio/gloo:0.5.6\nimagePullPolicy: Always\nname: gloo\nports:")
	)

	It("undefined version doesn't replace", func() {
		actual := install.UpdateBytesWithVersion(initial, "undefined")
		Expect(actual).To(BeEquivalentTo(initial))
	})

	It("defined version does replace", func() {
		actual := install.UpdateBytesWithVersion(initial, "0.5.6")
		Expect(actual).To(BeEquivalentTo(updated))
	})
})
