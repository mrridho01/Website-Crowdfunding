package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignId(campaignId uint) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

// generate instance struct repository agar bisa diakses dari package luar
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByCampaignId(campaignId uint) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Where("campaign_id = ?", campaignId).Preload("User").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
