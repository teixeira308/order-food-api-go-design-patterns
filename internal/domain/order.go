package domain

import (
	"time"
)

type Order interface {
	GetID() string
	GetStatus() OrderStatus
	GetCreatedAt() time.Time
	GetCustomerID() string
	GetItems() []OrderItem
	GetOrderType() string
}

type OrderStatus string

const (
	StatusReceived OrderStatus = "RECEIVED"
)

const (
	OrderTypeRegular = "regular"
	OrderTypeExpress = "express"
)

type OrderItem struct {
	ProductID string  `json:"product_id" bson:"product_id"`
	Name      string  `json:"name" bson:"name"`
	Price     float64 `json:"price" bson:"price"`
	Quantity  int     `json:"quantity" bson:"quantity"`
}

type RegularOrder struct {
	ID         string      `bson:"_id,omitempty"` // Mongo cria o ObjectID aqui
	CustomerID string      `bson:"customer_id"`
	Items      []OrderItem `bson:"items"`
	Status     OrderStatus `bson:"status"`
	CreatedAt  time.Time   `bson:"created_at"`
	OrderType  string      `bson:"order_type"`
}

func (o *RegularOrder) GetID() string           { return o.ID }
func (o *RegularOrder) GetStatus() OrderStatus  { return o.Status }
func (o *RegularOrder) GetCreatedAt() time.Time { return o.CreatedAt }
func (o *RegularOrder) GetCustomerID() string   { return o.CustomerID }
func (o *RegularOrder) GetItems() []OrderItem   { return o.Items }
func (o *RegularOrder) GetOrderType() string    { return o.OrderType }

type ExpressOrder struct {
	ID            string        `bson:"_id,omitempty"`
	CustomerID    string        `bson:"customer_id"`
	Items         []OrderItem   `bson:"items"`
	Status        OrderStatus   `bson:"status"`
	CreatedAt     time.Time     `bson:"created_at"`
	DeliveryLimit time.Duration `bson:"delivery_limit"`
	OrderType     string        `bson:"order_type"`
}

func (o *ExpressOrder) GetID() string           { return o.ID }
func (o *ExpressOrder) GetStatus() OrderStatus  { return o.Status }
func (o *ExpressOrder) GetCreatedAt() time.Time { return o.CreatedAt }
func (o *ExpressOrder) GetCustomerID() string   { return o.CustomerID }
func (o *ExpressOrder) GetItems() []OrderItem   { return o.Items }
func (o *ExpressOrder) GetOrderType() string    { return o.OrderType }
