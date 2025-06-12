package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	PackSize           int    `env:"COLLECTOR_PACK_SIZE,notEmpty"`
	NatsURL            string `env:"NATS_URL,notEmpty"`
	NatsGoodsTopic     string `env:"NATS_GOODS_TOPIC,notEmpty"`
	ClickhouseURL      string `env:"CLICKHOUSE_URL,notEmpty"`
	ClickhouseDB       string `env:"CLICKHOUSE_DB,notEmpty"`
	ClickhouseUser     string `env:"ROOT_USER,notEmpty"`
	ClickhousePassword string `env:"ROOT_PASSWORD,notEmpty"`
}

func Load() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return &cfg, err
}
