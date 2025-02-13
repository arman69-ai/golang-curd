package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Load ENV
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️  Gagal memuat .env file, gunakan default!")
	}
}

// Generate Token JWT
func GenerateToken(userID uint) (string, error) {
	LoadEnv()
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Verify Token JWT
func VerifyToken(tokenString string) (*jwt.Token, error) {
	LoadEnv()
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tidak valid")
		}
		return []byte(secret), nil
	})

	return token, err
}
