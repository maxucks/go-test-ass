package router

import (
	"database/sql"
	"net/http"
	com "test/internal/app/common"
	"test/internal/app/config"
	"test/internal/app/controllers"
	"test/internal/app/managers"
	"test/internal/app/repos"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

type healthResponse struct {
	message string
}

func Health(w http.ResponseWriter, r *http.Request) {
	com.JSON(w, healthResponse{
		message: "service is healthy",
	})
}

type Controllers struct {
	goods *controllers.GoodsController
}

func Setup(cfg *config.Config, db *sql.DB, nc *nats.Conn, rdb *redis.Client) *chi.Mux {
	repo := repos.NewGoods(db)
	cache := managers.NewRedisCache(rdb, managers.WithExpiration(cfg.CacheExpiration))
	pub := managers.NewNatsPublisher(nc, cfg.NatsGoodsTopic)
	controller := controllers.NewGoods(repo, pub, cache)

	r := chi.NewRouter()

	r.Get("/health", Health)

	r.Get("/projects/goods", controller.List)
	r.Post("/projects/{projectID}/goods", controller.Create)
	r.Patch("/projects/{projectID}/goods/{goodsID}", controller.Update)
	r.Delete("/projects/{projectID}/goods/{goodsID}", controller.Delete)
	r.Patch("/projects/{projectID}/goods/{goodsID}/priority", controller.Reprioritize)

	return r
}
