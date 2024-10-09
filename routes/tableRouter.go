package routes

import (
	contoller "restaurant-management/controller"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine){
    incomingRoutes.GET("/tables",contoller.GetTables())
	incomingRoutes.GET("/tables/:table_id",contoller.GetTable())
	incomingRoutes.POST("/tables"contoller.createMenu())
	incomingRoutes.PATCH("tables/:table_id",contoller.UpdateTable())
}