package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	CustomerName     string `json:"customer_name"`
	PickupLocation   string `json:"pickup_location"`
	DeliveryLocation string `json:"delivery_location"`
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/orders", createOrderHandler)

	fmt.Println("Order Service is running on port 8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Order Service is healthy")
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("New order received:")
	fmt.Println("Customer:", order.CustomerName)
	fmt.Println("Pickup:", order.PickupLocation)
	fmt.Println("Delivery:", order.DeliveryLocation)

	response := map[string]interface{}{
		"message": "Order created successfully",
		"order":   order,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
