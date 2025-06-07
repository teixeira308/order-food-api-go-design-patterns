package handler

import "github.com/gin-gonic/gin"

func RegisterOrderRoutes(r *gin.Engine, h *OrderHandler) {
	orders := r.Group("v1/orders")
	{
		orders.POST("/", h.CreateOrder)
		orders.GET("/:id", h.GetOrder)
		orders.GET("/", h.GetOrders)
	}
}
