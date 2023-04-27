package main

import (
	"log"
	"net/http"
	"startup-crowdfunding/auth"
	"startup-crowdfunding/campaign"
	"startup-crowdfunding/handler"
	"startup-crowdfunding/helper"
	"startup-crowdfunding/transaction"
	"startup-crowdfunding/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("entity/startup.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&user.User{}, &campaign.Campaign{}, &campaign.CampaignImage{}, &transaction.Transaction{})

	//create repository instance
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	//create service instance
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)

	//create handler instance
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	//static image routing
	router.Static("/image", "./images")
	api := router.Group("/api/v1")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/session", userHandler.Login)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/campaign", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.POST("/campaign-image", authMiddleware(authService, userService), campaignHandler.UploadCampaignImage)

	api.GET("/campaign", campaignHandler.GetCampaigns)
	api.GET("/campaign/:id", campaignHandler.GetCampaign)
	api.GET("campaign/:id/transaction", authMiddleware(authService, userService), transactionHandler.GetCampaignTransaction)
	api.GET("/transaction", authMiddleware(authService, userService), transactionHandler.GetUserTransaction)

	api.PUT("/campaign/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)

	router.Run()
}

// middleware untuk end point yang memerlukan authorization
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil nilai header Authorization : Bearer
		authHeader := c.GetHeader("Authorization")

		// cek apakah ada field Bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// abort request
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		// ambil nilai token
		arraySlice := strings.Split(authHeader, " ")
		if len(arraySlice) == 2 {
			tokenString = arraySlice[1]
		}

		// validasi token yang telah diambil
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// claim/ambil data berdasarkan token yang telah di validate
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := uint(claim["user_id"].(float64))
		// ambil user dari repo berdasarkan userId
		user, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set context user yang telah diambil
		c.Set("currentUser", user)
	}
}
