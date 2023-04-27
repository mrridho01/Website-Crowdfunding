package helper

import (
	"startup-crowdfunding/campaign"
	"startup-crowdfunding/transaction"
	"startup-crowdfunding/user"
	"strings"
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
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID               uint                           `json:"id"`
	Name             string                         `json:"name"`
	ShortDescription string                         `json:"short_description"`
	Description      string                         `json:"description"`
	ImageURL         string                         `json:"image_url"`
	GoalAmount       int                            `json:"goal_amount"`
	CurrentAmount    int                            `json:"current_amount"`
	UserId           uint                           `json:"user_id"`
	Slug             string                         `json:"slug"`
	Perks            []string                       `json:"perks"`
	User             CampaignDetailUserFormatter    `json:"user"`
	Images           []CampaignDetailImageFormatter `json:"images"`
}

type CampaignDetailUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignDetailImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type CampaignTransactionFormatter struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	Id        uint                                  `json:"id"`
	Amount    int                                   `json:"amount"`
	Status    string                                `json:"status"`
	CreatedAt time.Time                             `json:"created_at"`
	Campaign  UserTransactionCampaignImageFormatter `json:"campaign"`
}

type UserTransactionCampaignImageFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
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
		Slug:             campaign.Slug,
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

func FormatCampaignDetail(campaign campaign.Campaign) CampaignDetailFormatter {
	formatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		UserId:           campaign.UserId,
	}

	// field perks
	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	formatter.Perks = perks

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	} else {
		formatter.ImageURL = ""
	}

	// field user
	CampaignDetailUserFormatter := CampaignDetailUserFormatter{
		Name:     campaign.User.Name,
		ImageURL: campaign.User.AvatarFileName,
	}

	formatter.User = CampaignDetailUserFormatter

	// field images
	var images []CampaignDetailImageFormatter

	for _, image := range campaign.CampaignImages {
		var campaignDetailImageFormatter CampaignDetailImageFormatter

		campaignDetailImageFormatter.ImageURL = image.FileName
		if image.IsPrimary == 1 {
			campaignDetailImageFormatter.IsPrimary = true
		} else {
			campaignDetailImageFormatter.IsPrimary = false
		}

		images = append(images, campaignDetailImageFormatter)
	}

	formatter.Images = images

	return formatter
}

func FormatCampaignTransaction(transaction transaction.Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{
		Id:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter
}

func FormatCampaignTransactions(transactions []transaction.Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

func FormatUserTransaction(transaction transaction.Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{
		Id:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	campaignFormatter := UserTransactionCampaignImageFormatter{
		Name: transaction.Campaign.Name,
	}
	// cek apakah campaign memiliki image
	campaignFormatter.ImageURL = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []transaction.Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var userTransactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, formatter)
	}

	return userTransactionsFormatter
}

func FormatError(err error) []string {
	// array string untuk membungkus error validasi
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
