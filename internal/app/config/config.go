package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port            int    `env:"SRV_PORT" envDefault:"8000"`
	DatabaseURL     string `env:"PG_URL,notEmpty"`
	RedisPort       int    `env:"REDIS_PORT,notEmpty"`
	CacheExpiration int    `env:"CACHE_EXPIRATION,notEmpty"`
	NatsPort        int    `env:"NATS_PORT,notEmpty"`
	NatsGoodsTopic  string `env:"NATS_GOODS_TOPIC,notEmpty"`
}

func Load() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return &cfg, err
}
