package router

import (
	"github.com/gin-gonic/gin"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/handler"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/repository"
)

func SetupRouter(repo repository.OrderRepository) *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	orderHandler := handler.NewOrderHandler(repo)

	handler.RegisterOrderRoutes(r, orderHandler)
	return r
}
