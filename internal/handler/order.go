package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/domain"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/handler/dto"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/repository"
)

type OrderHandler struct {
	repo repository.OrderRepository
}

func NewOrderHandler(repo repository.OrderRepository) *OrderHandler {
	return &OrderHandler{repo: repo}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
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

	id, err := h.repo.Save(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao salvar pedido"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        id,
		"status":    string(order.GetStatus()),
		"createdAt": order.GetCreatedAt(),
	})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pedido não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         order.GetID(),
		"status":     order.GetStatus(),
		"created_at": order.GetCreatedAt(),
		"customerID": order.GetCustomerID(),
		"items":      order.GetItems(),
		"orderType":  order.GetOrderType(),
	})
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	orders, err := h.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar pedidos"})
		return
	}

	result := make([]gin.H, 0, len(orders))
	for _, order := range orders {
		result = append(result, gin.H{
			"id":         order.GetID(),
			"status":     order.GetStatus(),
			"created_at": order.GetCreatedAt(),
			"customerID": order.GetCustomerID(),
			"items":      order.GetItems(),
			"orderType":  order.GetOrderType(),
		})
	}

	c.JSON(http.StatusOK, result)
}
