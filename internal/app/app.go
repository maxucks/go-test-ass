package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"test/internal/app/config"
	"test/internal/app/router"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	db, err := sql.Open("postgres", cfg.DB.URL)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	natsURL := fmt.Sprintf("nats://localhost:%d", cfg.Nats.Port)
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("failed to establish nats connection: %s", err)
	}
	defer nc.Drain()

	r := router.Setup(cfg, db, nc, cfg.Nats.GoodsTopic)

	addr := fmt.Sprintf(":%d", cfg.Srv.Port)
	log.Printf("listening at %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %s", err)
		}
	}
}
