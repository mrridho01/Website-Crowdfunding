package transaction

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CampaignId uint
	UserId     uint
	Amount     int
	Status     string
	Code       string
}
