package config

import "github.com/kelseyhightower/envconfig"

type RuntimeConfig struct {
	Tag     string `envconfig:"TAGGED_VERSION"`
	Sha     string `envconfig:"SHA"`
	Version string `envconfig:"VERSION"`
}

func MustGetRuntimeConfig() *RuntimeConfig {
	var cfg RuntimeConfig
	envconfig.MustProcess("", &cfg)
	if cfg.Version == "" {
		if cfg.Tag == "" {
			cfg.Version = cfg.Sha
		} else {
			cfg.Version = cfg.Tag[1:]
		}
	}
	if cfg.Version == "" {
		cfg.Version = "undefined"
	}
	return &cfg
}
