package main

import (
	"golang/config"
	"golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.POST("/register", routes.RegisterUser)
	r.GET("/users", routes.GetUsers)
	r.POST("/login", routes.Login)



	r.Run(":4008") // http://localhost:4007 дээр ажиллана
}
