package model

type Order struct {
	CustomerName     string `json:"customer_name"`
	PickupLocation   string `json:"pickup_location"`
	DeliveryLocation string `json:"delivery_location"`
}
