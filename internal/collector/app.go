package collector

import (
	"context"
	"log"
	"test/internal/collector/config"
	"test/internal/collector/services"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/nats-io/nats.go"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatalf("failed to establish nats connection: %s", err)
	}
	defer nc.Drain()

	clickhouseConn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.ClickhouseURL},
		Auth: clickhouse.Auth{
			Database: cfg.ClickhouseDB,
			Username: cfg.ClickhouseUser,
			Password: cfg.ClickhousePassword,
		},
	})
	if err != nil {
		log.Fatalf("failed to establish clickhouse connection: %s", err)
	}

	ctx := context.Background()

	if err := clickhouseConn.Ping(ctx); err != nil {
		log.Fatalf("failed to ping clickhouse: %s", err)
	}

	service := services.NewCollector(clickhouseConn, nc, cfg.PackSize)
	if err := service.Run(ctx, cfg); err != nil {
		log.Fatalf("failed to init collector service: %s", err)
	}
	defer service.Close()

	log.Printf("collector service is running")
	select {}
}
