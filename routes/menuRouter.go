package routes

import (
controller "restaurant-management/contollers"
"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menus",contoller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id",controller.GetMenu())
	incomingRoutes.POST("/menus",contoller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id",controller.UpdateMenu())
}