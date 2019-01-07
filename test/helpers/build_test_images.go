package helpers

import (
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/solo-io/solo-kit/test/helpers"
	"hash/crc32"
	"os"
	"os/exec"
)

var glooComponents = []string{
	"gloo",
	"discovery",
	"gateway",
	"ingress",
}
var versionTag = ""

func Version() string {
	if versionTag != "" {
		return versionTag
	}
	tag := os.Getenv("VERSION")
	// if no tag set, default to a hash of the user's hostname
	if tag == "" {
		if host, err := os.Hostname(); err == nil {
			tag = hash(host)
		} else {
			tag = helpers.RandString(4)
		}
	}

	versionTag = "testing-" + tag
	return versionTag
}

func hash(h string) string {
	crc32q := crc32.MakeTable(0xD5828281)
	return fmt.Sprintf("%08x", crc32.Checksum([]byte(h), crc32q))
}

// builds and pushes all docker containers needed for test
func BuildPushContainers(push bool) error {
	if os.Getenv("SKIP_BUILD") == "1" {
		return nil
	}
	version := Version()
	os.Setenv("VERSION", version)

	// make the gloo containers
	for _, component := range []string{"gloo", "discovery", "kube-ingress-controller", "upstream-discovery"} {
		arg := component
		arg += "-docker"

		cmd := exec.Command("make", arg)
		cmd.Dir = GlooDir()
		cmd.Stdout = ginkgo.GinkgoWriter
		cmd.Stderr = ginkgo.GinkgoWriter
		cmd.Env = os.Environ()
		if err := cmd.Run(); err != nil {
			return err
		}

		if push {
			arg += "-push"

			cmd := exec.Command("make", arg)
			cmd.Dir = GlooDir()
			cmd.Stdout = ginkgo.GinkgoWriter
			cmd.Stderr = ginkgo.GinkgoWriter
			cmd.Env = os.Environ()
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	}

	// TODO (ilackarms): build test containers
	//for _, path := range []string{
	//	filepath.Join(GlooTestContainersDir(), "testrunner"),
	//	filepath.Join(KubeE2eDirectory(), "containers", "event-emitter"),
	//	filepath.Join(KubeE2eDirectory(), "containers", "upstream-for-events"),
	//	filepath.Join(KubeE2eDirectory(), "containers", "grpc-test-service"),
	//} {
	//	dockerOrg := os.Getenv("DOCKER_ORG")
	//	if dockerOrg == "" {
	//		dockerOrg = "soloio"
	//	}
	//	fullImage := dockerOrg + "/" + filepath.Base(path) + ":" + Version()
	//	log.Debugf("TEST: building fullImage %v", fullImage)
	//	cmd := exec.Command("make", "docker")
	//	cmd.Dir = path
	//	cmd.Stdout = ginkgo.GinkgoWriter
	//	cmd.Stderr = ginkgo.GinkgoWriter
	//	if err := cmd.Run(); err != nil {
	//		return err
	//	}
	//	if push {
	//		cmd = exec.Command("docker", "push", fullImage)
	//		cmd.Stdout = ginkgo.GinkgoWriter
	//		cmd.Stderr = ginkgo.GinkgoWriter
	//		if err := cmd.Run(); err != nil {
	//			return err
	//		}
	//	}
	//	cmd = exec.Command("make", "clean")
	//	cmd.Dir = path
	//	cmd.Stdout = ginkgo.GinkgoWriter
	//	cmd.Stderr = ginkgo.GinkgoWriter
	//	cmd.Run()
	//}
	return nil
}
