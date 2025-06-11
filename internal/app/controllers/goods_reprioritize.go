package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	com "test/internal/app/common"
	"test/internal/app/models"

	"github.com/go-chi/chi/v5"
)

type reprioritizeBody struct {
	NewPriority int `json:"newPriority"`
}

type reprioritizeResponse struct {
	Priorities []*models.ReprioritizedGoods `json:"priorities"`
}

func (c *GoodsController) Reprioritize(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	goodsID := chi.URLParam(r, "goodsID")

	// TODO: Validate
	// TODO: Check for existance

	fmt.Println(projectID, goodsID)

	var body reprioritizeBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		com.Error(w, err)
		return
	}

	com.JSON(w, reprioritizeResponse{
		Priorities: []*models.ReprioritizedGoods{},
	})
}
