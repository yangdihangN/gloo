package helpers

import (
	"bytes"
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

func DeployGlooWithHelm(namespace, imageVersion string) error {
	values, err := ioutil.TempFile("", "gloo-test-")
	if err != nil {
		return err
	}
	defer os.Remove(values.Name())
	if _, err := io.Copy(values, GlooHelmValues(imageVersion)); err != nil {
		return err
	}

	cmd := exec.Command("helm", "template", GlooHelmChartDir())
	cmd.Dir = GlooDir()
	cmd.Stdout = ginkgo.GinkgoWriter
	cmd.Stderr = ginkgo.GinkgoWriter
	cmd.Env = os.Environ()
	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrapf(err, "running command %v failed: %s", cmd.Args, out)
	}

	installBytes, err := exec.Command("helm", "template", HelmDirectory(),
		"--namespace", namespace,
		"-n", "test",
		"-f", values.Name()).CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "running helm template: %v", string(installBytes))
	}

	kubeApply := exec.Command("kubectl", "apply", "-f", "-")
	kubeApply.Stdin = bytes.NewBuffer(installBytes)
	cmd.Stdout = ginkgo.GinkgoWriter
	cmd.Stderr = ginkgo.GinkgoWriter
	cmd.Env = os.Environ()
	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrapf(err, "running command %v failed: %s", cmd.Args, out)
	}

	return nil
}

func GlooHelmValues(version string) io.Reader {
	b := &bytes.Buffer{}

	err := template.Must(template.ParseGlob(`
# note: these values must remain consistent with 
# install/helm/gloo/values.yaml
namespace:
  create: false
rbac:
  create: false

deployment:
  imagePullPolicy: IfNotPresent
  gloo:
    xdsPort: 9977
    image: soloio/gloo:{{ .Version }}
    replicas: 1
  discovery:
    image: soloio/discovery:{{ .Version }}
    replicas: 1
  gateway:
    image: soloio/gateway:{{ .Version }}
    replicas: 1
  gatewayProxy:
    image: soloio/gloo-envoy-wrapper:{{ .Version }}
    httpPort: 8080
    replicas: 1
  ingress:
    image: soloio/ingress:{{ .Version }}
    replicas: 1
  ingressProxy:
    image: soloio/gloo-envoy-wrapper:{{ .Version }}
    httpPort: 8080
    replicas: 1
`, )).Execute(b, struct {
		Version string
	}{
		Version: version,
	})
	if err != nil {
		panic(err)
	}

	return b
}

func WaitGlooPods() error {
	if err := WaitPodsRunning(glooComponents...); err != nil {
		return err
	}
	return nil
}

func WaitPodsRunning(podNames ...string) error {
	finished := func(output string) bool {
		return strings.Contains(output, "Running") || strings.Contains(output, "ContainerCreating")
	}
	for _, pod := range podNames {
		if err := WaitPodStatus(pod, "Running", finished); err != nil {
			return err
		}
	}
	return nil
}

func WaitPodsTerminated(podNames ...string) error {
	for _, pod := range podNames {
		finished := func(output string) bool {
			return !strings.Contains(output, pod)
		}
		if err := WaitPodStatus(pod, "terminated", finished); err != nil {
			return err
		}
	}
	return nil
}

func WaitPodStatus(pod, status string, finished func(output string) bool) error {
	timeout := time.Second * 20
	interval := time.Millisecond * 1000
	tick := time.Tick(interval)

	log.Debugf("waiting %v for pod %v to be %v...", timeout, pod, status)
	for {
		select {
		case <-time.After(timeout):
			return fmt.Errorf("timed out waiting for %v to be %v", pod, status)
		case <-tick:
			out, err := KubectlOut("get", "pod", "-l", "gloo="+pod)
			if err != nil {
				return fmt.Errorf("failed getting pod: %v", err)
			}
			if strings.Contains(out, "CrashLoopBackOff") {
				out = KubeLogs(pod)
				return errors.Errorf("%v in crash loop with logs %v", pod, out)
			}
			if strings.Contains(out, "ErrImagePull") || strings.Contains(out, "ImagePullBackOff") {
				out, _ = KubectlOut("describe", "pod", "-l", "gloo="+pod)
				return errors.Errorf("%v in ErrImagePull with description %v", pod, out)
			}
			if finished(out) {
				return nil
			}
		}
	}
}

func KubeLogs(pod string) string {
	out, err := KubectlOut("logs", "-l", "gloo="+pod)
	if err != nil {
		out = err.Error()
	}
	return out
}
