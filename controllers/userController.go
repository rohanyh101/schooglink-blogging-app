package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/schooglink_assignment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("user_id")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var user models.User
		err := UserCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if user.Username == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdStr := c.GetString("user_id")
		userId, err := primitive.ObjectIDFromHex(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateObj := bson.M{}

		if user.Username != nil {
			updateObj["username"] = user.Username
		}

		if user.Password != nil {
			hash := HashPassword(*user.Password)
			updateObj["password"] = hash
		}

		updateObj["updated_at"] = time.Now()

		filter := bson.M{"_id": bson.M{"$eq": userId}}
		update := bson.M{"$set": updateObj}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
	}
}
