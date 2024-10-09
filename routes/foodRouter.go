package routes

import (
	controller "restaurant-managment/controllers"
	"github.com/gin-gonic/gin"

)
func FoodRoutes(incomingRoutes *gin.Engine){
  incomingRoutes.GET("/foods",contoller.GetFoods())
  incomingRoutes.GET("/foods/:food_id",controller.GetFood())
  incomingRoutes.POST("/foods",controller.CreateFood())
  incomingRoutes.PATCH("/foods/:food_id",controller.UpdateFood())
}