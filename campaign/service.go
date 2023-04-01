package campaign

type Service interface {
	GetCampaigns(userId uint) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
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
