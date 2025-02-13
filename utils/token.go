package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-crud/config"
)

// Load .env saat package ini digunakan
func init() {
	config.LoadEnv()
}

// GenerateToken membuat JWT token
func GenerateToken(userID uint) (string, error) {
	// Ambil secret key dari .env
	secretKey := config.GetEnv("JWT_SECRET")

	// Buat claims token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expired dalam 24 jam
	}

	// Buat token dengan metode signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token dengan secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
