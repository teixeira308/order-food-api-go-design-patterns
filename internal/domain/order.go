package domain

import (
	"time"
)

type Order interface {
	GetID() string
	GetStatus() OrderStatus
	GetCreatedAt() time.Time
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
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

type RegularOrder struct {
	ID         string
	CustomerID string
	Items      []OrderItem
	Status     OrderStatus
	CreatedAt  time.Time
}

func (o *RegularOrder) GetID() string           { return o.ID }
func (o *RegularOrder) GetStatus() OrderStatus  { return o.Status }
func (o *RegularOrder) GetCreatedAt() time.Time { return o.CreatedAt }

type ExpressOrder struct {
	ID            string
	CustomerID    string
	Items         []OrderItem
	Status        OrderStatus
	CreatedAt     time.Time
	DeliveryLimit time.Duration
}

func (o *ExpressOrder) GetID() string           { return o.ID }
func (o *ExpressOrder) GetStatus() OrderStatus  { return o.Status }
func (o *ExpressOrder) GetCreatedAt() time.Time { return o.CreatedAt }
