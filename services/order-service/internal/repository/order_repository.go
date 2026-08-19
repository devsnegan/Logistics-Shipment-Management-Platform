package repository

import (
	"database/sql"

	"github.com/devsnegan/logistics-platform/order-service/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(order *model.Order) (*model.Order, error) {
	query := `
		INSERT INTO orders (
			customer_name,
			pickup_location,
			delivery_location
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := r.db.QueryRow(
		query,
		order.CustomerName,
		order.PickupLocation,
		order.DeliveryLocation,
	).Scan(
		&order.ID,
		&order.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetOrders() ([]*model.Order, error) {
	query := `
		SELECT
			id,
			customer_name,
			pickup_location,
			delivery_location,
			created_at
		FROM orders
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*model.Order

	for rows.Next() {
		var order model.Order

		err := rows.Scan(
			&order.ID,
			&order.CustomerName,
			&order.PickupLocation,
			&order.DeliveryLocation,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}