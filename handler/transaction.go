package handler

import (
	"net/http"
	"startup-crowdfunding/helper"
	"startup-crowdfunding/transaction"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		// format error validation failed
		error := helper.FormatError(err)
		errorMessage := gin.H{"error": error}

		response := helper.APIResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetTransactionsByCampaignId(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get campaign transactions", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)

}
