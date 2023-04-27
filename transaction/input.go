package transaction

import "startup-crowdfunding/user"

type GetCampaignTransactionDetailInput struct {
	Id   uint `uri:"id" binding:"required"`
	User user.User
}
