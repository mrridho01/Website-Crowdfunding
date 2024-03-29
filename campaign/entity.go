package campaign

import (
	"startup-crowdfunding/user"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	UserId           uint
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	gorm.Model
	CampaignId uint
	FileName   string
	IsPrimary  int
}
