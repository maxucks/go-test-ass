package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	com "test/internal/app/common"
	"test/internal/app/models"
)

type listResponse struct {
	Meta  models.PaginationMeta `json:"meta"`
	Goods []models.Goods        `json:"goods"`
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

	fmt.Println(limit, offset)

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
