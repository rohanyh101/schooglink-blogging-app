package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/schooglink_assignment/controllers"
	"github.com/roh4nyh/schooglink_assignment/middleware"
)

func BlogPostRoutes(incomingRoutes *gin.RouterGroup) {
	blogPostRoutes := incomingRoutes.Group("/posts")
	blogPostRoutes.Use(middleware.AuthenticateUser())

	// blog post crud
	blogPostRoutes.GET("/", controllers.GetBlogPosts())
	blogPostRoutes.POST("/", controllers.CreateBlogPost())
	blogPostRoutes.GET("/:post_id", controllers.GetBlogPost())
	blogPostRoutes.PUT("/:post_id", controllers.UpdateBlogPost())
	blogPostRoutes.DELETE("/:post_id", controllers.DeleteBlogPost())

	// get all blogposts of a single user
	incomingRoutes.GET("/user/posts", controllers.GetBlogPostsByUser())
}
