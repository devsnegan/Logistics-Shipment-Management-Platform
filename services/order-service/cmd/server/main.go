package main

import (
	"fmt"
	"net/http"

	"github.com/devsnegan/logistics-platform/order-service/internal/handler"
	"github.com/devsnegan/logistics-platform/order-service/internal/repository"
	"github.com/devsnegan/logistics-platform/order-service/internal/service"
)

func main() {
	repo := repository.NewOrderRepository()

	orderService := service.NewOrderService(repo)

	orderHandler := handler.NewOrderHandler(orderService)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Order Service is healthy")
	})

	fmt.Println("Order Service started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}

	_ = orderHandler
}
