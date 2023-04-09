package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId uint) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(id GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(inputImage CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// Service untuk mendapatkan list campaign
func (s *service) GetCampaigns(userId uint) ([]Campaign, error) {
	// apabila userId tidak kosong, ambil campaign berdasarkan userId
	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	//ambil semua campaign
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

// service get campaign by id
func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.Id)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		UserId:           input.User.ID,
	}

	// generate slug
	slugName := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugName)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil

}

func (s *service) UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputId.Id)
	if err != nil {
		return campaign, err
	}

	// cek authorization
	if campaign.UserId != inputData.User.ID {
		return campaign, errors.New("not authorized")
	}

	//update data campaign dari inputData. Slug by design tidak di update
	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(inputImage CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	isPrimary := 0
	if inputImage.IsPrimary {
		_, err := s.repository.MarkAllImageAsNonPrimary(inputImage.CampaignId)
		if err != nil {
			return CampaignImage{}, err
		}

		isPrimary = 1
	}

	campaignImage := CampaignImage{
		CampaignId: inputImage.CampaignId,
		IsPrimary:  isPrimary,
		FileName:   fileLocation,
	}

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
