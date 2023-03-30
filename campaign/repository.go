package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId uint) ([]Campaign, error)
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
	err := r.db.Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUSerId(userID uint) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
