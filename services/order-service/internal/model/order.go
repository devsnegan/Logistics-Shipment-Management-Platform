package model

import "time"

type Order struct {
	ID               int       `json:"id"`
	CustomerName     string    `json:"customer_name"`
	PickupLocation   string    `json:"pickup_location"`
	DeliveryLocation string    `json:"delivery_location"`
	CreatedAt        time.Time `json:"created_at"`
}