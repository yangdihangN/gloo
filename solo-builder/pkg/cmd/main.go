package main

import (
	"context"
	"fmt"
	"github.com/solo-io/gloo/solo-builder/pkg/config"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	ctx := context.TODO()
	artifacts := config.MustReadArtifactsConfig(ctx)
	runtime := config.MustGetRuntimeConfig()
	mustEnsureArtifactsDir(ctx)
	mustBuild(ctx, artifacts.Build, runtime)
	mustDocker(ctx, artifacts.Docker, runtime)
	mustHelm(ctx, artifacts.Helm, runtime)
}

const (
	ArtifactsDir = "_artifacts"
)

var (
	DefaultOs = []string {"linux"}
)

func mustEnsureArtifactsDir(ctx context.Context) {
	if _, err := os.Stat(ArtifactsDir); os.IsNotExist(err) {
		err = os.Mkdir(ArtifactsDir, os.ModePerm)
		if err != nil {
			contextutils.LoggerFrom(ctx).Fatalw("Failed to create artifacts directory", zap.Error(err))
		}
		return
	} else if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Could not check for artifacts directory", zap.Error(err))
	}
}

func mustBuild(ctx context.Context, build config.Build, runtime *config.RuntimeConfig) {
	mustGoBuild(ctx, build.Go, runtime)
}

func mustGoBuild(ctx context.Context, goBuild config.Go, runtime *config.RuntimeConfig) {
	// CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-X github.com/solo-io/gloo/pkg/version.Version=dev" -gcflags=all="-N -l" -o /Users/rick/code/src/github.com/solo-io/gloo/_output/gloo-linux-amd64 projects/gloo/cmd/main.go
	for _, binary := range goBuild.Binaries {
		if binary.Os == nil {
			binary.Os = DefaultOs
		}
		for _, osStr := range binary.Os {
			binaryName := fmt.Sprintf("%s-%s-%s", binary.Name, osStr, "amd64")
			parts := []string {
				"build",
				"-ldflags", fmt.Sprintf("-X %s=%s", goBuild.Version, runtime.Version),
				"-gcflags", goBuild.GcFlags,
				"-o", filepath.Join(ArtifactsDir, binaryName),
				binary.Entrypoint,
			}
			cmd := exec.Command("go", parts...)
			cmd.Env = append(os.Environ(), goEnv(osStr)...)
			fmt.Printf("Building %s\n", binaryName)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf(string(output))
				contextutils.LoggerFrom(ctx).Fatalw("Error creating artifact",
					zap.Error(err))
			}
		}
	}
}

func goEnv(os string) []string {
	return []string{
		"CGO_ENABLED=0",
		"GOOS=" + os,
		"GOARCH=amd64",
	}
}

func mustDocker(ctx context.Context, docker config.Docker, runtime *config.RuntimeConfig) {
	for _, registry := range docker.Registries {
		for _, container := range docker.Containers {
			dockerTag := fmt.Sprintf("%s:%s", filepath.Join(registry, container.Name), runtime.Version)
			cmd := exec.Command("docker", "build", "-t", dockerTag, "-f", container.Dockerfile, "_artifacts")
			fmt.Printf("Building docker container %s\n", dockerTag)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf(string(output))
				contextutils.LoggerFrom(ctx).Fatalw("Error building docker container",
					zap.Error(err))
			}

			cmd = exec.Command("docker", "push", dockerTag)
			fmt.Printf("Pushing docker container %s\n", dockerTag)
			output, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Printf(string(output))
				contextutils.LoggerFrom(ctx).Fatalw("Error pushing docker container",
					zap.Error(err))
			}
		}
	}
}

func mustHelm(ctx context.Context, helm config.Helm, runtime *config.RuntimeConfig) {
	for _, chart := range helm.Charts {
		mustHelmChart(ctx, chart, runtime)
	}
}

func mustHelmChart(ctx context.Context, chart config.Chart, runtime *config.RuntimeConfig) {
	generate := exec.Command("go", "run", chart.Generator, runtime.Version)
	fmt.Printf("Generating helm chart %s\n", chart.Name)
	output, err := generate.CombinedOutput()
	if err != nil {
		fmt.Printf(string(output))
		contextutils.LoggerFrom(ctx).Fatalw("Error generating helm chart",
			zap.Error(err))
	}
	pkg := exec.Command("helm", "package", "--destination", ArtifactsDir, chart.Directory)
	output, err = pkg.CombinedOutput()
	if err != nil {
		fmt.Printf(string(output))
		contextutils.LoggerFrom(ctx).Fatalw("Error packaging helm chart",
			zap.Error(err))
	}

	for _, manifest := range chart.Manifests {
		cmd := []string {"helm", "template",
			chart.Directory,
			"--namespace", "gloo-system",
			"--set", "namespace.create=true",
		}
		if manifest.Values != "" {
			valuesPath := filepath.Join(chart.Directory, manifest.Values)
			cmd = append(cmd, "--values", valuesPath)
		}

		fmt.Printf("Generating manifest %s\n", manifest.Name)
		manifestCmd := exec.Command(cmd[0], cmd[1:]...)
		output, err = manifestCmd.CombinedOutput()
		if err != nil {
			fmt.Printf(string(output))
			contextutils.LoggerFrom(ctx).Fatalw("Error generating manifest",
				zap.Error(err))
		}
		manifestPath := filepath.Join(ArtifactsDir, manifest.Name)
		err := ioutil.WriteFile(manifestPath, output, os.ModePerm)
		if err != nil {
			contextutils.LoggerFrom(ctx).Fatalw("Error writing manifest",
				zap.Error(err))
		}
	}
}
