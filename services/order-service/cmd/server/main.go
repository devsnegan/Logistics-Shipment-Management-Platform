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
http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		orderHandler.CreateOrder(w, r)

	case http.MethodGet:
		orderHandler.GetOrders(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
})
	fmt.Println("Order service running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
