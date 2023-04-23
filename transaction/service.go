package transaction

type Service interface {
	GetTransactionsByCampaignId(campaignId uint) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionsByCampaignId(campaignId uint) ([]Transaction, error) {
	transactions, err := s.repository.FindByCampaignId(campaignId)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
