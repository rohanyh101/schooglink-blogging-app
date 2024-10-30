package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/schooglink_assignment/controllers"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/auth/register", controllers.UserSignUp())
	incomingRoutes.POST("/auth/login", controllers.UserLogIn())
}
