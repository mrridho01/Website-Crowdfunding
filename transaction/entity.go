package transaction

import (
	"startup-crowdfunding/user"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CampaignId uint
	UserId     uint
	Amount     int
	Status     string
	Code       string
	User       user.User
}
