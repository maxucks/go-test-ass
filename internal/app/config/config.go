package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Srv Srv `envPrefix:"SRV_"`
}

type Srv struct {
	Port int `env:"PORT" envDefault:"8000"`
}

func Load() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return &cfg, err
}
