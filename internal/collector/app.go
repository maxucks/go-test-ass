package collector

import (
	"fmt"
	"log"
	"test/internal/collector/config"
	"test/internal/collector/services"

	"github.com/nats-io/nats.go"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	natsURL := fmt.Sprintf("nats://localhost:%d", cfg.NatsPort)
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("failed to establish nats connection: %s", err)
	}
	defer nc.Drain()

	service := services.NewCollector(nc, cfg.PackSize)
	if err := service.Init(cfg); err != nil {
		log.Fatalf("failed to init collector service: %s", err)
	}
	defer service.Close()

	log.Printf("collector service is running")
	select {}
}
