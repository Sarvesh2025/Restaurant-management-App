package routes

import (
	controller "restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users",controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
	incomingRoutes.POST("/users/signup",controller.SignUp())
	incomingRoutes.POST("/users/login",contoller.Login())
}