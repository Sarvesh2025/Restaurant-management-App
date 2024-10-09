package routes

import (
controller "restaurant-management/contollers"
"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders",contoller.GetOrders())
	incomingRoutes.GET("/orders/:order_id",controller.GetOrder())
	incomingRoutes.POST("/orders",contoller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id",controller.UpdateOrder())
}