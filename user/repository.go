package user

import "gorm.io/gorm"

// kontrak untuk struct user
type Repository interface {
	Save(user User) (User, error)
}

// Struct untuk menggunakan instance db yang telah dibuat di main.go
type repository struct {
	db *gorm.DB
}

// Membuat instance struct repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
