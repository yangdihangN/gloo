package kube2e_test

import (
	"os"
	"testing"
	"time"

	"github.com/solo-io/gloo/test/helpers"
	stringutils "github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	// TODO(ilackarms): tie testrunner to solo CI test containers and then handle image tagging
	defaultTestRunnerImage = "soloio/testrunner:latest"
)

func TestKube2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kube2e Suite")
}

var namespace string
var testRunnerPort int32
var _ = BeforeSuite(func() {
	if os.Getenv("RUN_KUBE2E_TESTS") != "1" {
		Skip("This test builds and deploys images to dockerhub and kubernetes, " +
			"and is disabled by default. To enable, set RUN_KUBE2E_TESTS=1 in your env.")
	}
	// todo (ilackarms): move randstring to stringutils package
	namespace = "a" + stringutils.RandString(8)
	testRunnerPort = 1234

	err := setup.SetupKubeForTest(namespace)
	Expect(err).NotTo(HaveOccurred())
	err = helpers.DeployTestRunner(namespace, defaultTestRunnerImage, testRunnerPort)
	Expect(err).NotTo(HaveOccurred())
	// build and push images for test
	version := helpers.TestVersion()
	err = helpers.BuildPushContainers(version, true, true)
	Expect(err).NotTo(HaveOccurred())
	err = helpers.DeployGlooWithHelm(namespace, version, true)
	Expect(err).NotTo(HaveOccurred())
	err = helpers.WaitGlooPods(time.Minute, time.Second)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	setup.TeardownKube(namespace)
})
