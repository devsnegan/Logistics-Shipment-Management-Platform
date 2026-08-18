package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devsnegan/logistics-platform/order-service/internal/database"
	"github.com/devsnegan/logistics-platform/order-service/internal/handler"
	"github.com/devsnegan/logistics-platform/order-service/internal/repository"
	"github.com/devsnegan/logistics-platform/order-service/internal/service"
)

func main() {
	// Connect to PostgreSQL
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer db.Close()

	fmt.Println("Database connected successfully")

	// Create repository
	orderRepository := repository.NewOrderRepository(db)

	// Create service
	orderService := service.NewOrderService(orderRepository)

	// Create handler
	orderHandler := handler.NewOrderHandler(orderService)

	// Register routes
	http.HandleFunc("/orders", orderHandler.CreateOrder)

	fmt.Println("Order service running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
