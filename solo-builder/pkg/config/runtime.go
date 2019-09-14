package config

import "github.com/kelseyhightower/envconfig"

type RuntimeConfig struct {
	Tag     string `envconfig:"TAGGED_VERSION"`
	Sha     string `envconfig:"SHA"`
	Version string
}

func MustGetRuntimeConfig() *RuntimeConfig {
	var cfg RuntimeConfig
	envconfig.MustProcess("", &cfg)
	if cfg.Tag == "" {
		cfg.Version = cfg.Sha
	} else {
		cfg.Version = cfg.Tag[1:]
	}
	return &cfg
}
