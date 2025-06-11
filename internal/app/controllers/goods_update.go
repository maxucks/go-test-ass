package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	com "test/internal/app/common"
	"test/internal/app/models"

	"github.com/go-chi/chi/v5"
)

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

	fmt.Println(projectID, goodsID)

	// TODO: Check for existance

	var res *models.Goods

	com.JSON(w, res)
}
