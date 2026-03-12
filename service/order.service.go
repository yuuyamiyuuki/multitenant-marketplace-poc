package service

import (
	"errors"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/repository"

	"github.com/google/uuid"
)

type OrderService interface {
	Checkout(input dto.CheckoutInput, tenantID, userID uuid.UUID) (*repository.Order, error)
	ConcludeOrder(orderID uuid.UUID) error
	ListMyOrders(tenantID, userID uuid.UUID) ([]*repository.Order, error)
}

type orderService struct {
	uow repository.UnitOfWork
}

func NewOrderService(uow repository.UnitOfWork) OrderService {
	return &orderService{uow: uow}
}

func (s *orderService) ListMyOrders(tenantID, userID uuid.UUID) ([]*repository.Order, error) {
	clientRepo := s.uow.ClientRepo()
	client, err := clientRepo.FindByUserIDAndTenant(userID, tenantID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, middleware.NewNotFound("client profile not found for user", err)
		}
		return nil, middleware.NewInternal("failed to fetch client profile", err)
	}

	orders, err := s.uow.OrderRepo().FindByClientIDAndTenant(client.ID, tenantID)
	if err != nil {
		return nil, middleware.NewInternal("failed to fetch orders", err)
	}

	return orders, nil
}

func (s *orderService) Checkout(input dto.CheckoutInput, tenantID, userID uuid.UUID) (*repository.Order, error) {
	var createdOrder *repository.Order

	err := s.uow.Do(func(uow repository.UnitOfWork) error {
		clientRepo := uow.ClientRepo()
		client, err := clientRepo.FindByUserIDAndTenant(userID, tenantID)
		if err != nil {
			if err.Error() == "record not found" {
				return middleware.NewNotFound("client profile not found for user", err)
			}
			return middleware.NewInternal("failed to fetch client profile", err)
		}

		cartRepo := uow.CartRepo()
		cart, err := cartRepo.FindByUserIDAndTenant(userID, tenantID)
		if err != nil {
			if err.Error() == "record not found" {
				return middleware.NewNotFound("cart not found", err)
			}
			return middleware.NewInternal("failed to retrieve cart", err)
		}

		if len(cart.Items) == 0 {
			return middleware.NewBadRequest("cart is empty", errors.New("cannot checkout empty cart"))
		}

		productRepo := uow.ProductRepo()

		var totalAmount float64
		orderItems := make([]repository.OrderItem, 0, len(cart.Items))

		for _, item := range cart.Items {
			product, err := productRepo.FindByIDAndTenant(item.ProductID, tenantID)
			if err != nil {
				return middleware.NewInternal("failed to fetch product in cart", err)
			}

			if product.Stock < item.Quantity {
				return middleware.NewBadRequest("insufficient stock for product "+product.Name, errors.New("insufficient stock"))
			}

			product.Stock -= item.Quantity
			if err := productRepo.Update(product); err != nil {
				return middleware.NewInternal("failed to update product stock", err)
			}

			subtotal := float64(item.Quantity) * product.Price
			totalAmount += subtotal

			orderItems = append(orderItems, repository.OrderItem{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Amount:    subtotal,
			})
		}

		newOrder := &repository.Order{
			TenantID:      tenantID,
			ClientID:      client.ID,
			TotalAmount:   totalAmount,
			Items:         orderItems,
			PaymentMethod: input.PaymentMethod,
			Status:        "waiting_payment", // Initial state
		}

		if err := uow.OrderRepo().Create(newOrder); err != nil {
			return middleware.NewInternal("failed to create checkout record", err)
		}

		if err := cartRepo.ClearItems(cart.ID, tenantID); err != nil {
			return middleware.NewInternal("failed to clear cart", err)
		}

		createdOrder = newOrder
		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}

func (s *orderService) ConcludeOrder(orderID uuid.UUID) error {
	orderRepo := s.uow.OrderRepo()
	order, err := orderRepo.FindByID(orderID)
	if err != nil {
		if err.Error() == "record not found" {
			return middleware.NewNotFound("order not found", err)
		}
		return middleware.NewInternal("failed to fetch order", err)
	}

	if order.Status != "waiting_payment" {
		return middleware.NewBadRequest("order is not waiting for payment", errors.New("invalid status transition"))
	}

	order.Status = "concluded"
	if err := orderRepo.Update(order); err != nil {
		return middleware.NewInternal("failed to update order status", err)
	}

	return nil
}
