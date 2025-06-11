package controllers

import (
	"fmt"
	"net/http"
	com "test/internal/app/common"
	"test/internal/app/models"

	"github.com/go-chi/chi/v5"
)

type createBody struct {
	Name string `json:"name"`
}

type createResponse struct {
	Name string `json:"name"`
}

func (c *GoodsController) Create(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")

	fmt.Println(projectID)

	var goods *models.Goods

	com.JSON(w, goods)
}
