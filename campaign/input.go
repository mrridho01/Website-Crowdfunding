package campaign

type GetCampaignDetailInput struct {
	Id uint `uri:"id" binding:"required"`
}
