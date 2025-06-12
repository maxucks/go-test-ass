package controllers

import (
	"log"
	"net/http"
	"strconv"
	com "test/internal/app/common"
	"test/internal/app/models"
)

type listResponse struct {
	Meta  listResponseMeta `json:"meta"`
	Goods []*models.Goods  `json:"goods"`
}

type listResponseMeta struct {
	Total   int `json:"total"`
	Removed int `json:"removed"`
	Offset  int `json:"offset"`
	Limit   int `json:"limit"`
}

func (c *GoodsController) List(w http.ResponseWriter, r *http.Request) {
	var limit, offset int = 10, 0
	var err error

	rawLimit := r.URL.Query().Get("limit")
	if rawLimit != "" {
		limit, err = strconv.Atoi(rawLimit)
		if err != nil {
			com.BadRequest(w, com.WithDetails("limit is not a number"))
			return
		}
	}

	rawOffset := r.URL.Query().Get("offset")
	if rawOffset != "" {
		offset, err = strconv.Atoi(rawOffset)
		if err != nil {
			com.BadRequest(w, com.WithDetails("offset is not a number"))
			return
		}
	}

	ctx := r.Context()

	meta, err := c.cache.GetGoodsMetadata(ctx)
	if err != nil {
		log.Printf("err reading cache: %s\n", err)
	}
	if meta == nil {
		log.Println("No cache")
		total, removed, err := c.repo.GetPaginationMeta(r.Context())
		if err != nil {
			com.Error(w, err)
			return
		}
		meta = &models.PaginationMeta{
			Total:   total,
			Removed: removed,
		}
		if err := c.cache.CacheGoodsMetadata(ctx, *meta); err != nil {
			log.Printf("err caching meta: %s\n", err)
		}
	}

	goods, err := c.cache.GetGoods(ctx, offset, limit)
	if err != nil {
		log.Printf("err reading cache: %s\n", err)
	}
	if goods == nil {
		goods, err = c.repo.Get(ctx, limit, offset)
		if err != nil {
			com.Error(w, err)
			return
		}
		c.cache.CacheGoods(ctx, offset, limit, goods)
	}

	com.JSON(w, listResponse{
		Meta: listResponseMeta{
			Total:   meta.Total,
			Removed: meta.Removed,
			Limit:   limit,
			Offset:  offset,
		},
		Goods: goods,
	})
}
