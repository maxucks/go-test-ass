package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Srv  Srv      `envPrefix:"SRV_"`
	DB   Database `envPrefix:"PG_"`
	Nats Nats     `envPrefix:"NATS_"`
}

type Srv struct {
	Port int `env:"PORT" envDefault:"8000"`
}

type Database struct {
	URL string `env:"URL,notEmpty"`
}

type Nats struct {
	Port       int    `env:"PORT,notEmpty"`
	GoodsTopic string `env:"GOODS_TOPIC,notEmpty"`
}

func Load() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return &cfg, err
}
