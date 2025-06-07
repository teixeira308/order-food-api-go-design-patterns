package main

import (
	"log"

	"github.com/teixeira308/order-food-api-go-design-patterns/internal/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Servidor rodando em http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
