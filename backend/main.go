package main

import (
	"go-backend/config"
	"go-backend/models"
	"go-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB холбох
	config.ConnectDB()

	// Auto migrate
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// User register route
	r.POST("/register", routes.RegisterUser)

	r.Run(":8080") // localhost:8080 дээр сервер асаах
}
