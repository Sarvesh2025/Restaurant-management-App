package routes

import (
controller "restaurant-management/contollers"
"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices",contoller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id",controller.GetInvoice())
	incomingRoutes.POST("/invoices",contoller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id",controller.UpdateInvoice())
}