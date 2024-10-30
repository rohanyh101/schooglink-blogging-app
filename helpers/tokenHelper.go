package helpers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/roh4nyh/schooglink_assignment/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedUserDetails struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

const (
	userDatabaseName   = "Cluster0"
	userCollectionName = "users"
)

var UserCollection *mongo.Collection = database.OpenCollection(userDatabaseName, userCollectionName)
var USER_SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateUserToken(userId, username, email string) (signedToken string, err error) {
	claims := &SignedUserDetails{
		UserID:   userId,
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(USER_SECRET_KEY))
	if err != nil {
		msg := fmt.Sprintf("Error signing Token: %v", err)
		return "", errors.New(msg)
	}

	return token, nil
}

func UpdateUserToken(signedToken, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updatedAt})
	upsert := true

	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := UserCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)

	if err != nil {
		log.Panic(err)
		return
	}
}

func ValidateUserToken(signedToken string) (claims *SignedUserDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedUserDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(USER_SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedUserDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("the token has expired")
		msg = err.Error()
		return
	}
	return claims, msg
}
