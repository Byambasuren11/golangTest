package routes

import (
	"context"
	"golang/config"
	"golang/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)
func GetUsers(c *gin.Context) {
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(500, gin.H{"error": "Error decoding user"})
			return
		}
		user.Password = ""
		users = append(users, user)
	}

	c.JSON(200, gin.H{"users": users})
}