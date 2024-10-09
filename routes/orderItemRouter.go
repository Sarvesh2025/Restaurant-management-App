package routes

import (
	contoller "restaurant-management/controller"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine){
    incomingRoutes.GET("/orderItems",contoller.GetorderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id",contoller.GetorderItem())
	incomingRoutes.GET("/orderItems-order/:order_id",controller.GetorderItemsByOrder())
	incomingRoutes.POST("/orderItems"contoller.createOrderItem())
	incomingRoutes.PATCH("orderItems/:orderItem_id",contoller.UpdateOrderItem())
}