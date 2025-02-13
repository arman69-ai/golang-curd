package controllers

import (
	"net/http"

	"github.com/arman69-ai/golang-crud/database"
	"github.com/arman69-ai/golang-crud/helpers"
	"github.com/arman69-ai/golang-crud/models"

	"github.com/gin-gonic/gin"

	"github.com/arman69-ai/golang-crud/utils"
)

// Register User
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password sebelum disimpan
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hash password"})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil"})
}

// Login User
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	// Validasi password
	if !user.CheckPassword(input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	// Generate JWT
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// LoginHandler contoh handler login
func LoginHandler(c *gin.Context) {
	// Contoh user ID (biasanya didapat dari database)
	userID := uint(1)

	// Generate token
	token, err := utils.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
