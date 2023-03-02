package user

import "gorm.io/gorm"

// kontrak untuk struct user
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
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

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
