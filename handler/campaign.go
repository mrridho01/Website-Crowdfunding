package handler

import (
	"fmt"
	"net/http"
	"startup-crowdfunding/campaign"
	"startup-crowdfunding/helper"
	"startup-crowdfunding/user"
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

	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", helper.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	//tangkap parameter id dari url / uri
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		// format error validation failed
		error := helper.FormatError(err)
		errorMessage := gin.H{"error": error}

		response := helper.APIResponse("Failed to get detail campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get detail campaign", http.StatusOK, "success", helper.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorValidationMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Fail create campaign", http.StatusUnprocessableEntity, "error", errorValidationMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Fail create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success create campaign", http.StatusOK, "success", helper.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	//tangkap parameter id dari url / uri
	var inputId campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		// format error validation failed
		error := helper.FormatError(err)
		errorMessage := gin.H{"error": error}

		response := helper.APIResponse("Fail to update campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//tangkap input dari parameter
	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatError(err)
		errorValidationMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Fail to update campaign", http.StatusUnprocessableEntity, "error", errorValidationMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputId, inputData)
	if err != nil {
		response := helper.APIResponse("Fail to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update campaign", http.StatusOK, "success", helper.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadCampaignImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed upload campaign image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	campaignImage, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed upload campaign image", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// ambil user yang sedang upload data
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userId := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userId, campaignImage.Filename)

	// save image ke folder path
	err = c.SaveUploadedFile(campaignImage, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// save campaign image ke db
	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success upload campaign image", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
