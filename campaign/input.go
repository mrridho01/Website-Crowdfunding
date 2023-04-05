package campaign

import "startup-crowdfunding/user"

type GetCampaignDetailInput struct {
	Id uint `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	GoalAmount       int    `json:"goal_amount"`
	Perks            string `json:"perks"`
	User             user.User
}
