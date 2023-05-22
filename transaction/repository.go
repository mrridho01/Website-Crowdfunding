package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignId(campaignId uint) ([]Transaction, error)
	FindByUserId(userId uint) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
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
	err := r.db.Where("campaign_id = ?", campaignId).Order("id desc").Preload("User").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindByUserId(userId uint) ([]Transaction, error) {
	var transaction []Transaction
	// load campaign images yang tidak relasi langsung dengan transaksi
	err := r.db.Where("user_id = ?", userId).Order("id desc").Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
