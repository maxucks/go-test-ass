package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	com "test/internal/app/common"
	"test/internal/app/models"

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

	fmt.Println(projectID, body.Name)

	var goods *models.Goods

	com.JSON(w, goods)
}
