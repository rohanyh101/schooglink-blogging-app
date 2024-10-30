package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/roh4nyh/schooglink_assignment/database"
	"github.com/roh4nyh/schooglink_assignment/helpers"
	"github.com/roh4nyh/schooglink_assignment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserCollectionName = "users"
)

var userValidate = validator.New()
var UserCollection *mongo.Collection = database.OpenCollection(DatabaseName, UserCollectionName)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword, foundUserPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(foundUserPassword), []byte(userPassword))
	check := true

	if err != nil {
		check = false
	}

	return check
}

func UserSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := userValidate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user already exists, please login"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()

		token, _ := helpers.GenerateUserToken(user.UserID, *user.Username, *user.Email)
		user.Token = &token

		resultInsertionNumber, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created. Error: %v", insertErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusCreated, resultInsertionNumber)
	}
}

func UserLogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid := VerifyPassword(*user.Password, *foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		if foundUser.Username == nil || foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		token, err := helpers.GenerateUserToken(foundUser.UserID, *foundUser.Username, *foundUser.Email)
		if err != nil || token == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		helpers.UpdateUserToken(token, foundUser.UserID)

		err = UserCollection.FindOne(ctx, bson.M{"_id": foundUser.ID}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser)
	}
}
