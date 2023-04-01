package helper

import (
	"startup-crowdfunding/campaign"
	"startup-crowdfunding/user"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserFormatter struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation"`
	Email      string    `json:"email"`
	Token      string    `json:"token"`
	Role       string    `json:"role"`
	ImageURL   string    `json:"image_url"`
}

type CampaignFormatter struct {
	ID               uint   `json:"id"`
	UserId           uint   `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatUser(user user.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		Role:       user.Role,
		ImageURL:   user.AvatarFileName,
	}

	return formatter
}

func FormatCampaign(campaign campaign.Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
	}
	//cek apakah campaign memiliki campaignImages
	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	} else {
		formatter.ImageURL = ""
	}

	return formatter
}

// formatter untuk slice Campaign
func FormatCampaigns(campaigns []campaign.Campaign) []CampaignFormatter {
	var campaignsFormatter []CampaignFormatter = []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatError(err error) []string {
	// array string untuk membungkus error validasi
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
