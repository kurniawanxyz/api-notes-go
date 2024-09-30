package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kurniawanxyz/crud-notes-go/config"
)

func GenerateJWT(data any) (string, error) {
	// Buat klaim (claims) untuk token
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token akan expire dalam 1 jam

	// Buat token dengan metode signing HMAC menggunakan klaim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Convert the secret key to a byte slice
	secretKey := []byte(config.ENV.JWTSecret)

	// Tanda tangani token dengan secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}