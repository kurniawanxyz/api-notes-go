package helper

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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


func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Ambil token dari header Authorization
        tokenString := c.GetHeader("Authorization")

        // Token harus dimulai dengan "Bearer "
        if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			HandleResponse(c, http.StatusUnauthorized, "Authorization header is required")
            c.Abort()
            return
        }

        // Hilangkan prefix "Bearer " untuk mendapatkan token murni
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Parse token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method")
            }
            return config.ENV.JWTSecret, nil
        })

        if err != nil || !token.Valid {
			HandleResponse(c, http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Jika token valid, lanjutkan ke handler berikutnya
        c.Next()
    }
}