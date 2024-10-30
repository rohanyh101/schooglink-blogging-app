package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/schooglink_assignment/database"
	"github.com/roh4nyh/schooglink_assignment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DatabaseName           = "Cluster0"
	BlogPostCollectionName = "blogposts"
)

var BlogPostCollection *mongo.Collection = database.OpenCollection(DatabaseName, BlogPostCollectionName)

func GetBlogPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var blogPosts []models.Post

		cursor, err := BlogPostCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing blog posts"})
			return
		}

		if err = cursor.All(ctx, &blogPosts); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing blog posts"})
			return
		}

		if len(blogPosts) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "no blog post available"})
			return
		}

		c.JSON(http.StatusOK, blogPosts)
	}
}

func GetBlogPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("post_id")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var blogPost models.Post
		err := BlogPostCollection.FindOne(ctx, bson.M{"post_id": postId}).Decode(&blogPost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "blog post not found"})
			return
		}

		if blogPost.Title == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "blog post not found"})
			return
		}

		c.JSON(http.StatusOK, blogPost)
	}
}

func CreateBlogPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var blogPost models.Post
		userIdStr := c.GetString("user_id")
		username := c.GetString("username")

		if err := c.BindJSON(&blogPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		blogPost.CreatedAt = time.Now()
		blogPost.UpdatedAt = time.Now()
		blogPost.ID = primitive.NewObjectID()
		blogPost.PostID = blogPost.ID.Hex()
		blogPost.Author = &username

		userId, err := primitive.ObjectIDFromHex(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
			return
		}
		blogPost.UserID = userId

		resultInsertionNumber, insertErr := BlogPostCollection.InsertOne(ctx, blogPost)
		if insertErr != nil {
			msg := "Blog post item was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusCreated, resultInsertionNumber)
	}
}

func UpdateBlogPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("post_id")
		userId := c.GetString("user_id")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var blogPost models.Post

		if err := c.BindJSON(&blogPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// check is bolg post exists and belongs to the user
		var existingPost models.Post
		err := BlogPostCollection.FindOne(ctx, bson.M{"post_id": postId}).Decode(&existingPost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "blog post not found"})
			return
		}

		if existingPost.UserID.Hex() != userId {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
			return
		}

		updateObj := bson.M{}

		if blogPost.Title != nil {
			updateObj["title"] = blogPost.Title
		}

		if blogPost.Content != nil {
			updateObj["content"] = blogPost.Content
		}

		if blogPost.Description != nil {
			updateObj["description"] = blogPost.Description
		}

		updateObj["updated_at"] = time.Now()

		filter := bson.M{"post_id": bson.M{"$eq": postId}}
		update := bson.M{"$set": updateObj}

		_, err = BlogPostCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating blog post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "blog post updated successfully"})
	}
}

func DeleteBlogPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("post_id")
		userId := c.GetString("user_id")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		// check is bolg post exists and belongs to the user
		var existingPost models.Post
		err := BlogPostCollection.FindOne(ctx, bson.M{"post_id": postId}).Decode(&existingPost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "blog post not found"})
			return
		}

		// Check if the user is the owner of the post
		if existingPost.UserID.Hex() != userId {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
			return
		}

		filter := bson.M{"post_id": postId}
		_, err = BlogPostCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting blog post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "blog post deleted successfully"})
	}
}

func GetBlogPostsByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdStr := c.GetString("user_id")
		userId, err := primitive.ObjectIDFromHex(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var blogPosts []models.Post

		cursor, err := BlogPostCollection.Find(ctx, bson.M{"user_id": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing blog posts"})
			return
		}

		if err = cursor.All(ctx, &blogPosts); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing blog posts"})
			return
		}

		if len(blogPosts) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "no blog post available"})
			return
		}

		c.JSON(http.StatusOK, blogPosts)
	}
}
