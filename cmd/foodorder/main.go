package main

import (
	"context"
	"log"
	"time"

	"github.com/teixeira308/order-food-api-go-design-patterns/internal/repository"
	"github.com/teixeira308/order-food-api-go-design-patterns/internal/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Erro conectando no MongoDB: %v", err)
	}

	repo := repository.NewMongoOrderRepository(client, "orderdb", "orders")

	r := router.SetupRouter(repo)

	log.Println("Servidor rodando em http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
