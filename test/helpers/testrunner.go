package helpers

import (
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

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
			// needed for WaitForPodsRunning
			Labels: map[string]string{"gloo": "testrunner"},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Image: image,
					Name:  "testrunner",
				},
			},
		},
	}); err != nil {
		return err
	}
	if err := WaitPodsRunning(time.Minute, time.Second, "testrunner"); err != nil {
		return err
	}
	return nil
}
