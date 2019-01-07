package kube2e_test

import (
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	stringutils "github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"testing"

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

var _ = BeforeSuite(func() {
	// todo (ilackarms): move randstring to stringutils package
	namespace = stringutils.RandString(8)

	err := setup.SetupKubeForTest(namespace)
	Expect(err).NotTo(HaveOccurred())
	// build and push images for test
	err = helpers.BuildPushContainers(true)
	Expect(err).NotTo(HaveOccurred())
	err = DeployTestRunner(namespace, defaultTestRunnerImage)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	err := setup.TeardownKube(namespace)
	Expect(err).NotTo(HaveOccurred())
})

func DeployTestRunner(namespace, image string) error {
	cfg, err := kubeutils.GetConfig("", "")
	if err != nil {
		return err
	}
	kube, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return err
	}
	if _, err := kube.CoreV1().Pods(namespace).Create(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "testrunner",
			Namespace: namespace,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Image: image,
				},
			},
		},
	}); err != nil {
		return err
	}
	if err := helpers.WaitPodsRunning("testrunner"); err != nil {
		return err
	}
	return nil
}
