package router

import (
	"net/http"
	"test/internal/app/config"
	"test/internal/app/controllers"

	"github.com/go-chi/chi/v5"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type Controllers struct {
	goods *controllers.GoodsController
}

func Setup(cfg *config.Config) *chi.Mux {
	ctrs := Controllers{
		goods: controllers.NewGoods(),
	}

	r := chi.NewRouter()

	r.Get("/health", Health)

	r.Get("/projects/goods", ctrs.goods.List)
	r.Post("/projects/{projectID}/goods", ctrs.goods.Create)
	r.Patch("/projects/{projectID}/goods/{goodsID}", ctrs.goods.Update)
	r.Delete("/projects/{projectID}/goods/{goodsID}", ctrs.goods.Delete)
	r.Patch("/projects/{projectID}/goods/{goodsID}/priority", ctrs.goods.Reprioritize)

	return r
}
