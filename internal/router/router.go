package router

import (
	"github.com/gin-gonic/gin"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Healthcheck
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Rotas da API
	handler.RegisterOrderRoutes(r)

	return r
}
