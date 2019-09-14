package config

import (
	"context"
	"github.com/ghodss/yaml"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
	"io/ioutil"
)

const (
	path = "artifacts.yaml"
)

func MustReadArtifactsConfig(ctx context.Context) Artifacts {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Could not read artifacts config", zap.Error(err))
	}

	var artifacts Artifacts
	if err := yaml.Unmarshal(bytes, &artifacts); err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Could not parse artifacts config", zap.Error(err))
	}

	return artifacts
}

type Artifacts struct {
	Build  Build  `json:"build,omitempty"`
	Docker Docker `json:"docker,omitempty"`
	Helm   Helm   `json:"helm,omitempty"`
}

type Build struct {
	Go Go `json:"go,omitempty"`
}

type Go struct {
	Version  string   `json:"version,omitempty"`
	GcFlags  string   `json:"gcFlags,omitempty"`
	Binaries []Binary `json:"binaries,omitempty"`
}

type Binary struct {
	Name       string   `json:"name,omitempty"`
	Os         []string `json:"os,omitempty"`
	Entrypoint string   `json:"entrypoint,omitempty"`
}

type Docker struct {
	Registries []string    `json:"registries,omitempty"`
	Containers []Container `json:"containers,omitempty"`
}

type Container struct {
	Name       string `json:"name,omitempty"`
	Dockerfile string `json:"dockerfile,omitempty"`
}

type Helm struct {
	Charts []Chart `json:"charts,omitempty"`
}

type Chart struct {
	Name      string     `json:"name,omitempty"`
	Directory string     `json:"directory,omitempty"`
	Generator string     `json:"generator,omitempty"`
	Manifests []Manifest `json:"manifests,omitempty"`
}

type Manifest struct {
	Name   string `json:"name,omitempty"`
	Values string `json:"values,omitempty"`
}
