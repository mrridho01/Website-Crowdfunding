package transaction

import (
	"errors"
	"startup-crowdfunding/campaign"
)

type Service interface {
	GetTransactionsByCampaignId(input GetCampaignTransactionDetailInput) ([]Transaction, error)
	GetTransactionsByUserId(userId uint) ([]Transaction, error) //userid didapat dari jwt token
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignId(input GetCampaignTransactionDetailInput) ([]Transaction, error) {
	// cek auth campaign transcation dengan user id yang bersesuaian
	campaign, err := s.campaignRepository.FindById(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("not owner of the campaign")
	}

	transactions, err := s.repository.FindByCampaignId(input.Id)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserId(userId uint) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{
		CampaignId: input.CampaignId,
		Amount:     input.Amount,
		UserId:     input.User.ID,
		Status:     "pending",
	}

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
