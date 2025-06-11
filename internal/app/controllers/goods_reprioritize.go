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

type reprioritizeBody struct {
	NewPriority int `json:"newPriority"`
}

type reprioritizeResponse struct {
	Priorities []*models.ReprioritizedGoods `json:"priorities"`
}

func (c *GoodsController) Reprioritize(w http.ResponseWriter, r *http.Request) {
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

	var body reprioritizeBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		com.Error(w, err)
		return
	}

	fmt.Println(projectID, goodsID, body.NewPriority)

	// TODO: Check for existance

	com.JSON(w, reprioritizeResponse{
		Priorities: []*models.ReprioritizedGoods{},
	})
}
