package service

import (
	"github.com/google/uuid"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/repository"
)

type CartService interface {
	UpsertCart(input dto.UpsertCartInput, tenantID uuid.UUID, userID uuid.UUID) (*repository.Cart, error)
	//DeleteCart(tenantID uuid.UUID) (*repository.Cart, error)
	GetCart(tenantID uuid.UUID, userID uuid.UUID) (*repository.Cart, error)
}

type cartService struct {
	cartRepo       repository.CartRepository
	productService ProductService
}

func NewCartService(repo repository.CartRepository, pSvc ProductService) CartService {
	return &cartService{
		cartRepo: repo, productService: pSvc,
	}
}

func (s *cartService) UpsertCart(input dto.UpsertCartInput, tenantID uuid.UUID, userID uuid.UUID) (*repository.Cart, error) {
	newCart := &repository.Cart{
		UserID:   userID,
		TenantID: tenantID,
		Products: input.Products,
	}
	err := s.cartRepo.Upsert(newCart)
	if err != nil {
		return nil, middleware.NewInternal("failed to create cart in the database", err)
	}

	return newCart, nil
}

func (s *cartService) GetCart(tenantID uuid.UUID, userID uuid.UUID) (*repository.Cart, error) {

	// find cart by user id and tenant
	// get products from list
	return nil, nil
}
