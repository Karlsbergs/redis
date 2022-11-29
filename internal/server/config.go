package server

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host string `envconfig:"HOST"`
	Port string `envconfig:"PORT"`
}

func ReadConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
