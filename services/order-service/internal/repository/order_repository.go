package repository

import "context"

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Create(ctx context.Context) {
	// Database logic will come here later.
}
