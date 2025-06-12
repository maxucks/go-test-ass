package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	com "test/internal/app/common"

	"github.com/go-chi/chi/v5"
)

type createBody struct {
	Name string `json:"name"`
}

func (c *GoodsController) Create(w http.ResponseWriter, r *http.Request) {
	rawProjectID := chi.URLParam(r, "projectID")
	projectID, err := strconv.Atoi(rawProjectID)
	if err != nil {
		com.BadRequest(w, com.WithDetails("projectID is not a number"))
		return
	}

	var body createBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		com.Error(w, err)
		return
	}

	goods, err := c.repo.Create(r.Context(), projectID, body.Name)
	if err != nil {
		com.Error(w, err)
		return
	}

	c.pub.PublishGoods(goods)

	com.JSON(w, goods)
}
