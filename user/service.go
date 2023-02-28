package user

import "golang.org/x/crypto/bcrypt"

// mewakili business logic aplikasi
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

//struct internal untuk mengakses interface Repository
type service struct {
	repository Repository
}

//fungsi untuk return instance dari struct service
func NewService(repository Repository) *service {
	return &service{repository}
}

//register user, dengan memanggil layer repository
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
