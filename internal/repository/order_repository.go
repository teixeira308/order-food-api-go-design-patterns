package repository

import (
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/domain"
)

// AQUI: Letra maiúscula para exportar
type OrderRepository interface {
	Save(order domain.Order) (string, error) // retorna o ID gerado
	FindByID(id string) (domain.Order, error)
	List() ([]domain.Order, error)
}
