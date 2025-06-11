package controllers

import (
	"fmt"
	"net/http"
	com "test/internal/app/common"
	"test/internal/app/models"
)

type listResponse struct {
	Meta  models.PaginationMeta `json:"meta"`
	Goods []models.Goods        `json:"goods"`
}

func (c *GoodsController) List(w http.ResponseWriter, r *http.Request) {
	rawLimit := r.URL.Query().Get("limit")
	rawOffset := r.URL.Query().Get("offset")

	// TODO: validate limit and offset

	fmt.Println(rawLimit, rawOffset)

	com.JSON(w, listResponse{
		Meta: models.PaginationMeta{
			Total:   0,
			Removed: 0,
			Limit:   10,
			Offset:  0,
		},
		Goods: []models.Goods{},
	})
}
