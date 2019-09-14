package main

import (
	"context"
	"fmt"
	"github.com/solo-io/gloo/solo-builder/pkg/config"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
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
