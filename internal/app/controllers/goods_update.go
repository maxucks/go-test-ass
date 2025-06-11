package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	com "test/internal/app/common"

	"github.com/go-chi/chi/v5"
)

type updateBody struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (c *GoodsController) Update(w http.ResponseWriter, r *http.Request) {
	rawProjectID := chi.URLParam(r, "projectID")
	projectID, err := strconv.Atoi(rawProjectID)
	if err != nil {
		com.BadRequest(w, com.WithDetails("projectID is not a number"))
		return
	}

	rawGoodsID := chi.URLParam(r, "goodsID")
	goodsID, err := strconv.Atoi(rawGoodsID)
	if err != nil {
		com.BadRequest(w, com.WithDetails("goodsID is not a number"))
		return
	}

	var body updateBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		com.Error(w, err)
		return
	}

	if body.Name == "" {
		com.BadRequest(w, com.WithDetails("name must not be empty"))
		return
	}

	ctx := r.Context()

	exists, err := c.repo.Exists(ctx, goodsID, projectID)
	if err != nil {
		com.Error(w, err)
		return
	}
	if !exists {
		com.NotFound(w)
		return
	}

	goods, err := c.repo.Update(r.Context(), goodsID, projectID, body.Name, body.Description)
	if err != nil {
		com.Error(w, err)
		return
	}

	com.JSON(w, goods)
}
