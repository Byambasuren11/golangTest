package routes

import (
	"go-backend/config"
	"go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register user
func RegisterUser(c *gin.Context) {
	var input models.User

	// Bind JSON → struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	input.Password = string(hashedPassword)

	// Save DB
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       input.ID,
			"username": input.Username,
			"email":    input.Email,
		},
	})
}
