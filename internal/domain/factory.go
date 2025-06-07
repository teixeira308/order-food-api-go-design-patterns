package domain

import (
	"errors"
	"time"
)

func NewOrder(orderType, customerID string, items []OrderItem, opts ...interface{}) (Order, error) {
	if customerID == "" {
		return nil, errors.New("customerID obrigatório")
	}
	if len(items) == 0 {
		return nil, errors.New("o pedido deve conter pelo menos um item")
	}

	switch orderType {
	case OrderTypeRegular:
		return &RegularOrder{
			ID:         "", // Mongo gera o ID
			CustomerID: customerID,
			Items:      items,
			Status:     StatusReceived,
			CreatedAt:  time.Now(),
			OrderType:  OrderTypeRegular,
		}, nil

	case OrderTypeExpress:
		if len(opts) == 0 {
			return nil, errors.New("delivery limit obrigatório para pedido express")
		}
		deliveryLimit, ok := opts[0].(time.Duration)
		if !ok {
			return nil, errors.New("delivery limit deve ser time.Duration")
		}
		return &ExpressOrder{
			ID:            "", // Mongo gera o ID
			CustomerID:    customerID,
			Items:         items,
			Status:        StatusReceived,
			CreatedAt:     time.Now(),
			DeliveryLimit: deliveryLimit,
			OrderType:     OrderTypeExpress,
		}, nil

	default:
		return nil, errors.New("tipo de pedido inválido")
	}
}
