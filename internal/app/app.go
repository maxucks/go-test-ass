package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"test/internal/app/config"
	"test/internal/app/router"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	r := router.Setup(cfg)

	addr := fmt.Sprintf(":%d", cfg.Srv.Port)
	if err := http.ListenAndServe(addr, r); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %s", err)
		}
	}
}
