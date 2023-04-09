package campaign

import "startup-crowdfunding/user"

type GetCampaignDetailInput struct {
	Id uint `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CreateCampaignImageInput struct {
	CampaignId uint `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"` //binding non required supaya bisa menerima nilai boolean false
	User       user.User
}
