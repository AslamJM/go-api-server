package utils

import (
	"kstrategy-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID uint `json:"id"`
	jwt.StandardClaims
}

func GenerateAccessToken(userId int) (string, error) {
	expirationTime := time.Now().Add(config.AccessTokenExpiry)
	claims := &Claims{
		UserID: uint(userId),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(userId uint) (string, error) {
	expirationTime := time.Now().Add(config.RefreshTokenExpiry)

	claims := &Claims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.JwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
