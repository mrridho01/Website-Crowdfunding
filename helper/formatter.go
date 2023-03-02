package helper

import (
	"startup-crowdfunding/user"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserFormatter struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation"`
	Email      string    `json:"email"`
	Token      string    `json:"token"`
	Role       string    `json:"role"`
}

func FormatUser(user user.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		Role:       user.Role,
	}

	return formatter
}

func FormatError(err error) []string {
	// array string untuk membungkus error validasi
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
