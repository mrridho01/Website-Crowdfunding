package handler

import (
	"net/http"
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

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, user)
}
