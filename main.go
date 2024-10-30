package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/roh4nyh/schooglink_assignment/routes"
)

func init() {
	if os.Getenv("ENV") == "production" {
		log.Printf("using .env production variables")
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		log.Printf("using .env development variables")
	}
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	gin.SetMode(gin.ReleaseMode)

	app := gin.New()
	app.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	app.Use(cors.New(config))

	apiV1 := app.Group("/api/v1")

	apiV1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "server is up and running..."})
	})

	routes.AuthRoutes(apiV1)

	routes.UserRoutes(apiV1)

	routes.BlogPostRoutes(apiV1)

	app.Run(fmt.Sprintf(":%s", PORT))
}
