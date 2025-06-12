package controllers

import (
	"net/http"
	"strconv"
	com "test/internal/app/common"
	"test/internal/app/models"

	"github.com/go-chi/chi/v5"
)

func (c *GoodsController) Delete(w http.ResponseWriter, r *http.Request) {
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

	removed, err := c.repo.Remove(r.Context(), goodsID, projectID)
	if err != nil {
		com.Error(w, err)
		return
	}

	c.pub.PublishGoods(removed)

	com.JSON(w, models.ShortGoods{
		Id:        removed.Id,
		ProjectId: removed.ProjectId,
		Removed:   removed.Removed,
	})
}
