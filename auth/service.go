package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID uint) (string, error)
}

type jwtService struct{}

//fake secret key
var SECRET_KEY = []byte("STARTUP_crowdfund1ng")

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
