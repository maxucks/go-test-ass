package router

import (
	"database/sql"
	"net/http"
	"test/internal/app/config"
	"test/internal/app/controllers"
	"test/internal/app/repos"

	"github.com/go-chi/chi/v5"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type Controllers struct {
	goods *controllers.GoodsController
}

func Setup(cfg *config.Config, db *sql.DB) *chi.Mux {
	repo := repos.NewGoods(db)
	controller := controllers.NewGoods(repo)

	r := chi.NewRouter()

	r.Get("/health", Health)

	r.Get("/projects/goods", controller.List)
	r.Post("/projects/{projectID}/goods", controller.Create)
	r.Patch("/projects/{projectID}/goods/{goodsID}", controller.Update)
	r.Delete("/projects/{projectID}/goods/{goodsID}", controller.Delete)
	r.Patch("/projects/{projectID}/goods/{goodsID}/priority", controller.Reprioritize)

	return r
}
