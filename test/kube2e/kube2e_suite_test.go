package kube2e_test

import (
	"github.com/solo-io/gloo/test/helpers"
	stringutils "github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"os"
	"testing"
	"time"

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
	// todo (ilackarms): move randstring to stringutils package
	namespace = stringutils.RandString(8)
	testRunnerPort = 1234

	err := setup.SetupKubeForTest(namespace)
	Expect(err).NotTo(HaveOccurred())
	err = helpers.DeployTestRunner(namespace, defaultTestRunnerImage, testRunnerPort)
	Expect(err).NotTo(HaveOccurred())
	// build and push images for test
	version := helpers.TestVersion()
	if os.Getenv("BUILD") == "1" {
		err = helpers.BuildPushContainers(version, true, true)
		Expect(err).NotTo(HaveOccurred())
	}
	err = helpers.DeployGlooWithHelm(namespace, version, true)
	Expect(err).NotTo(HaveOccurred())
	err = helpers.WaitGlooPods(time.Minute, time.Second)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	err := setup.TeardownKube(namespace)
	Expect(err).NotTo(HaveOccurred())
})
