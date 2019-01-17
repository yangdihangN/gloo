package kube2e_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"io/ioutil"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"path/filepath"
	"time"
)

var _ = Describe("Kube2e: Knative-Ingress", func() {
	BeforeEach(func() {
		deployKnative()
	})
	AfterEach(func() {
		deleteKnative()
	})
	It("works", func() {
		clusterIngressProxy := "clusteringress-proxy"
		clusterIngressPort := 80
		setup.CurlEventuallyShouldRespond(setup.CurlOpts{
			Protocol: "http",
			Path:     "/",
			Method:   "GET",
			Host:     "helloworld-go.default.example.com",
			Service:  clusterIngressProxy,
			Port:     clusterIngressPort,
		}, "Hello Go Sample v1!" , time.Minute * 2)
	})
})

func deployKnative() {
	b, err := ioutil.ReadFile(KnativeManifest())
	Expect(err).NotTo(HaveOccurred())

	err = helpers.RunCommandInput(string(b), true, "kubectl", "apply", "-f", "-")
	Expect(err).NotTo(HaveOccurred())
}

func deleteKnative() {
	b, err := ioutil.ReadFile(KnativeManifest())
	Expect(err).NotTo(HaveOccurred())

	err = helpers.RunCommandInput(string(b), true, "kubectl", "apply", "-f", "-")
	Expect(err).NotTo(HaveOccurred())
}

func KnativeManifest() string {
	return filepath.Join(helpers.GlooDir(), "test", "kube2e", "artifacts", "knative-no-istio.yaml")
}
