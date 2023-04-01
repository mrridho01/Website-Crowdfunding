package handler

import (
	"net/http"
	"startup-crowdfunding/campaign"
	"startup-crowdfunding/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// endpoint api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// tangkap parameter di url
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(uint(userId))
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
