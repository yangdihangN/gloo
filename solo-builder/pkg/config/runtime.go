package config

import (
	"github.com/kelseyhightower/envconfig"
	"os"
)

type RuntimeConfig struct {
	Version string `envconfig:"VERSION"`
}

func MustGetRuntimeConfig() *RuntimeConfig {
	var cfg RuntimeConfig
	envconfig.MustProcess("", &cfg)
	if cfg.Version == "" {
		cfg.Version = os.Getenv("TAGGED_VERSION")
		if cfg.Version != "" {
			cfg.Version = cfg.Version[1:]
		}
	}
	if cfg.Version == "" {
		cfg.Version = os.Getenv("SHA")
	}
	if cfg.Version == "" {
		cfg.Version = "undefined"
	}
	return &cfg
}
