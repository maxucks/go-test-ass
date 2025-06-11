package controllers

import "test/internal/app/ports"

type GoodsController struct {
	repo ports.GoodsRepo
}

func NewGoods(repo ports.GoodsRepo) *GoodsController {
	return &GoodsController{repo}
}
