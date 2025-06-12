package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	PackSize       int    `env:"COLLECTOR_PACK_SIZE,notEmpty"`
	NatsPort       int    `env:"NATS_PORT,notEmpty"`
	NatsGoodsTopic string `env:"NATS_GOODS_TOPIC,notEmpty"`
}

func Load() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return &cfg, err
}
