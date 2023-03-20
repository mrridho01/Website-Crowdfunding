package handler

import (
	"fmt"
	"net/http"
	"startup-crowdfunding/auth"
	"startup-crowdfunding/helper"
	"startup-crowdfunding/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

// membuat instance struct userHandler
func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// handler untuk endpoint register user
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
		return // return, so code below not executed
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Registering account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		response := helper.APIResponse("Registering account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helper.FormatUser(user, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// handler untuk endpoint login
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		// mapping errors ke dalam field errors
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helper.FormatUser(loggedUser, token)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

// handler untuk integrasi mendapatkan user yang sedang login
func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	formatter := helper.FormatUser(currentUser, "token")

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// error gagal validasi
		errors := helper.FormatError(err)
		// mapping errors ke dalam field errors
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// service untuk mengecek availability email
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	// menentukan message api response berdasarkan nilai isEmailAvailable
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	// apabila tidak ada error validasi dan error service yang terjadi
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// menangkap input dari user bukan dalam bentuk JSON, namun dalam form data
	file, err := c.FormFile("avatar")
	if err != nil {
		// helper bawaan gin untuk membuat map
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	// get current user berdasarkan context currentUser
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	_, err = h.userService.SaveAvatar(uint(userId), path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success upload avatar image", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
