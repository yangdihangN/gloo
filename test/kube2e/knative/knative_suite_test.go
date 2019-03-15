package knative_test

import (
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/testutils/clusterlock"
	"k8s.io/client-go/kubernetes"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/testutils"
	"github.com/solo-io/go-utils/testutils/helper"
	skhelpers "github.com/solo-io/solo-kit/test/helpers"
)

func TestKnative(t *testing.T) {
	if testutils.AreTestsDisabled() {
		return
	}
	skhelpers.RegisterCommonFailHandlers()
	skhelpers.SetupLog()
	RunSpecs(t, "Knative Suite")
}

var testHelper *helper.SoloTestHelper
var locker *clusterlock.TestClusterLocker

func MustKubeClient() kubernetes.Interface {
	restConfig, err := kubeutils.GetConfig("", "")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return kubeClient
}

var _ = BeforeSuite(func() {
	cwd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())
	testHelper, err = helper.NewSoloTestHelper(func(defaults helper.TestConfig) helper.TestConfig {
		defaults.RootDir = filepath.Join(cwd, "../../..")
		defaults.HelmChartName = "gloo"
		return defaults
	})
	Expect(err).NotTo(HaveOccurred())


	locker, err = clusterlock.NewTestClusterLocker(MustKubeClient(), testHelper.InstallNamespace)
	Expect(err).NotTo(HaveOccurred())
	Expect(locker.AcquireLock()).NotTo(HaveOccurred())
	// Install Gloo
	err = testHelper.InstallGloo(helper.KNATIVE, 5*time.Minute)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	defer locker.ReleaseLock()
	err := testHelper.UninstallGloo()
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() error {
		return testutils.Kubectl("get", "namespace", testHelper.InstallNamespace)
	}, "60s", "1s").Should(HaveOccurred())
})
