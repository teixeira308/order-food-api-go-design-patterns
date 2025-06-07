package dto

import (
	"time"

	"github.com/teixeira308/order-food-api-go-design-patterns/internal/domain"
)

type CreateOrderRequest struct {
	OrderType     string             `json:"order_type" binding:"required"`
	CustomerID    string             `json:"customer_id" binding:"required"`
	Items         []domain.OrderItem `json:"items" binding:"required,dive"`
	DeliveryLimit int64              `json:"delivery_limit"` // segundos, opcional
}

type CreateOrderResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
