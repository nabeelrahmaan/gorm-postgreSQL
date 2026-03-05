package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AccessSecret = []byte("access-secret")
var RefreshSecret = []byte("refresh-secret")

func GenerateAccessToken(userID uint) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(AccessSecret)
}

func GenerateRefreshToken(userID uint) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(RefreshSecret)
}