package service

import (
	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/repository"

	"github.com/google/uuid"
)

type OrderItemService interface {
	CreateOrderItem(input dto.CreateOrderItemInput, tenantID uuid.UUID) (*repository.OrderItem, error)
	GetOrderItem(id uuid.UUID, tenantID uuid.UUID) (*repository.OrderItem, error)
}

type orderItemService struct {
	orderitemRepo repository.OrderItemRepository
}

func NewOrderItemService(repo repository.OrderItemRepository) OrderItemService {
	return &orderItemService{
		orderitemRepo: repo,
	}
}

func (s *orderItemService) CreateOrderItem(input dto.CreateOrderItemInput, tenantID uuid.UUID) (*repository.OrderItem, error) {
	newOrderItem := &repository.OrderItem{
		Quantity: input.Quantity,
		Amount:   input.Amount,
		// missing arguments are handled by the caller or dto
	}

	err := s.orderitemRepo.Create(newOrderItem)
	if err != nil {
		return nil, middleware.NewInternal("failed to create a Order Item in database", err)
	}

	return newOrderItem, nil
}

func (s *orderItemService) GetOrderItem(id uuid.UUID, tenantID uuid.UUID) (*repository.OrderItem, error) {
	orderItem, err := s.orderitemRepo.FindByIDAndTenant(id, tenantID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, middleware.NewNotFound("the requested order item does not exist", err)
		}
		return nil, middleware.NewInternal("failed to fetch order item from database", err)
	}
	return orderItem, nil
}
