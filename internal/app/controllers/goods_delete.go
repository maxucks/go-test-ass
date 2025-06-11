package controllers

import (
	"fmt"
	"net/http"
	com "test/internal/app/common"

	"github.com/go-chi/chi/v5"
)

type deleteResponse struct {
	Id         int  `json:"id"`
	CampaignId int  `json:"campaignId"`
	Removed    bool `json:"removed"`
}

func (c *GoodsController) Delete(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	goodsID := chi.URLParam(r, "goodsID")

	// TODO: Validate
	// TODO: Check for existance

	fmt.Println(projectID, goodsID)

	com.JSON(w, deleteResponse{
		Id:         1,
		CampaignId: 0,
		Removed:    true,
	})
}
