package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	com "test/internal/app/common"

	"github.com/go-chi/chi/v5"
)

type deleteResponse struct {
	Id         int  `json:"id"`
	CampaignId int  `json:"campaignId"`
	Removed    bool `json:"removed"`
}

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

	fmt.Println(projectID, goodsID)

	// TODO: Check for existance

	com.JSON(w, deleteResponse{
		Id:         1,
		CampaignId: 0,
		Removed:    true,
	})
}
