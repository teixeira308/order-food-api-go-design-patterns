package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/domain"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/handler/dto"
)

func createOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order domain.Order
	var err error

	if req.OrderType == domain.OrderTypeExpress {
		order, err = domain.NewOrder(req.OrderType, req.CustomerID, req.Items, time.Duration(req.DeliveryLimit)*time.Second)
	} else {
		order, err = domain.NewOrder(domain.OrderTypeRegular, req.CustomerID, req.Items)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.CreateOrderResponse{
		ID:        order.GetID(),
		Status:    string(order.GetStatus()),
		CreatedAt: order.GetCreatedAt(),
	}

	c.JSON(http.StatusCreated, res)
}

func RegisterOrderRoutes(r *gin.Engine) {
	orders := r.Group("v1/orders")
	{
		orders.POST("/", createOrder)
		orders.GET("/:id", getOrder)
	}
}

func getOrder(c *gin.Context) {
	id := c.Param("id")
	// TODO: implementação da busca do pedido
	c.JSON(200, gin.H{"id": id, "status": "em preparo"})
}
