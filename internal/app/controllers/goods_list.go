package controllers

import (
	"net/http"
	"strconv"
	com "test/internal/app/common"
	"test/internal/app/models"
)

type listResponse struct {
	Meta  models.PaginationMeta `json:"meta"`
	Goods []*models.Goods       `json:"goods"`
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

	total, removed, err := c.repo.GetPaginationMeta(r.Context())
	if err != nil {
		com.Error(w, err)
		return
	}

	goods, err := c.repo.Get(r.Context(), limit, offset)
	if err != nil {
		com.Error(w, err)
		return
	}

	com.JSON(w, listResponse{
		Meta: models.PaginationMeta{
			Total:   total,
			Removed: removed,
			Limit:   limit,
			Offset:  offset,
		},
		Goods: goods,
	})
}
