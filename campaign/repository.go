package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId uint) ([]Campaign, error)
	FindById(Id uint) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImageAsNonPrimary(campaignId uint) (bool, error)
}

type repository struct {
	db *gorm.DB
}

// generate instance struct repository agar bisa diakses dari package luar
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// repo untuk get semua campaign tanpa filter
func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserId(userId uint) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindById(Id uint) (Campaign, error) {
	var campaign Campaign
	err := r.db.Where("id = ?", Id).Preload("CampaignImages").Preload("User").Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

/*Mengubah nilai is_primary images di database yang sebelumnya 1 menjadi 0,
apabila user ingin mengganti images baru menjadi is_primary*/
func (r *repository) MarkAllImageAsNonPrimary(campaignId uint) (bool, error) {
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignId).Update("is_primary", 0).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
