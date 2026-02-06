package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userId" : userID,
		"exp" : time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte(j.Secret))
}