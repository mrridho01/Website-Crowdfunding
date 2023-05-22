package transaction

import "startup-crowdfunding/user"

type GetCampaignTransactionDetailInput struct {
	Id   uint `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int  `json:"amount" binding:"required"`
	CampaignId uint `json:"campaign_id" binding:"required"`
	User       user.User
}
