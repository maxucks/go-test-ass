package controllers

import (
	"fmt"
	"net/http"
	com "test/internal/app/common"
	"test/internal/app/models"

	"github.com/go-chi/chi/v5"
)

// goodsID + projectID - check for existance
func (c *GoodsController) Update(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	goodsID := chi.URLParam(r, "goodsID")

	// TODO: Validate
	// TODO: Check for existance

	fmt.Println(projectID, goodsID)

	var res *models.Goods

	com.JSON(w, res)
}
