package transaction

import "gorm.io/gorm"

type Repository interface{}

type repository struct {
	db *gorm.DB
}
