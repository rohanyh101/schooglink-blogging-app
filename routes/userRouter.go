package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/schooglink_assignment/controllers"
	"github.com/roh4nyh/schooglink_assignment/middleware"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.Use(middleware.AuthenticateUser())

	// user crud
	incomingRoutes.GET("/users/profile/", controllers.GetUser())
	incomingRoutes.PUT("/users/profile/", controllers.UpdateUser())
}
