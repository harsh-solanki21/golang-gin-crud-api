package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessTokenSecret  = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	refreshTokenSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID, role string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessTokenSecret)
}

func GenerateRefreshToken(userID string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshTokenSecret)
}

func ValidateAccessToken(tokenString string) (*Claims, error) {
	return validateToken(tokenString, accessTokenSecret)
}

func ValidateRefreshToken(tokenString string) (*Claims, error) {
	return validateToken(tokenString, refreshTokenSecret)
}

func validateToken(tokenString string, secret []byte) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	return GenerateAccessToken(claims.UserID, claims.Role)
}
