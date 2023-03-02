package handler

import (
	"net/http"
	"startup-crowdfunding/helper"
	"startup-crowdfunding/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

// membuat instance struct userHandler
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// capture input dari user
	var input user.RegisterUserInput

	// error validasi akan muncul disini
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		// mapping errors ke dalam field errors
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Registering account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		errors := helper.FormatError(err)
		// mapping errors ke dalam field errors
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Registering account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helper.FormatUser(user, "token")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
