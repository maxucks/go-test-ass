package controllers

import (
	"test/internal/app/ports"
)

type GoodsController struct {
	repo ports.GoodsRepo
	pub  ports.QuePublisher
}

func NewGoods(repo ports.GoodsRepo, pub ports.QuePublisher) *GoodsController {
	return &GoodsController{repo, pub}
}
