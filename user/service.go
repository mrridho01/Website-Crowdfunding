package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// mewakili business logic aplikasi
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

// struct internal untuk mengakses interface Repository
type service struct {
	repository Repository
}

// fungsi untuk return instance dari struct service
func NewService(repository Repository) *service {
	return &service{repository}
}

// register user, dengan memanggil layer repository
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// service untuk login
func (s *service) Login(input LoginInput) (User, error) {
	// dapatkan email dan password dari input user
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		// custom error dengan built in errors package
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
