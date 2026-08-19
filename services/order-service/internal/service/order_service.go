package service

import (
	"github.com/devsnegan/logistics-platform/order-service/internal/model"
	"github.com/devsnegan/logistics-platform/order-service/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(order *model.Order) (*model.Order, error) {
	return s.repo.CreateOrder(order)
}

func (s *OrderService) GetOrders() ([]*model.Order, error) {
	return s.repo.GetOrders()
}

