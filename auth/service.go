package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{}

// fake secret key
var SECRET_KEY = []byte("STARTUP_crowdfund1ng")

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (j *jwtService) GenerateToken(userId uint) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//setelah generate token, untuk security maka perlu secret key untuk 'tanda tangan'
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (j *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	//parse token
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		//validasi apakah token nya HMAC
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		//kembalikan secret key, untuk di cek apakah dari secret key yang sama yang telah dibuat
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
