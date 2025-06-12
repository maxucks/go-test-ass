package controllers

import (
	"test/internal/app/ports"
)

type GoodsController struct {
	repo  ports.GoodsRepo
	pub   ports.QuePublisher
	cache ports.Cache
}

func NewGoods(repo ports.GoodsRepo, pub ports.QuePublisher, cache ports.Cache) *GoodsController {
	return &GoodsController{repo, pub, cache}
}
