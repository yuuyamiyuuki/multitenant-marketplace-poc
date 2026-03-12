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
	GetCart(tenantID uuid.UUID, userID uuid.UUID) (*dto.CartResponse, error)
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
	}

	for _, item := range input.Items {
		newCart.Items = append(newCart.Items, repository.CartItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	err := s.cartRepo.Upsert(newCart)
	if err != nil {
		return nil, middleware.NewInternal("failed to create cart in the database", err)
	}

	return newCart, nil
}

func (s *cartService) GetCart(tenantID uuid.UUID, userID uuid.UUID) (*dto.CartResponse, error) {
	cart, err := s.cartRepo.FindByUserIDAndTenant(userID, tenantID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, middleware.NewNotFound("cart not found", err)
		}
		return nil, middleware.NewInternal("failed to fetch cart", err)
	}

	productIDs := make([]uuid.UUID, 0, len(cart.Items))
	for _, item := range cart.Items {
		productIDs = append(productIDs, item.ProductID)
	}

	var products []*repository.Product
	if len(productIDs) > 0 {
		products, err = s.productService.GetProductsByIDs(productIDs, tenantID)
		if err != nil {
			return nil, err
		}
	}

	productMap := make(map[uuid.UUID]*repository.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	response := &dto.CartResponse{
		ID:         cart.ID,
		Items:      make([]dto.CartItemResponse, 0, len(cart.Items)),
		TotalItems: 0,
		TotalPrice: 0,
	}

	for _, item := range cart.Items {
		product, ok := productMap[item.ProductID]
		if !ok {
			continue
		}
		subtotal := float64(item.Quantity) * product.Price
		response.TotalPrice += subtotal
		response.TotalItems += item.Quantity
		response.Items = append(response.Items, dto.CartItemResponse{
			ProductID: product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  item.Quantity,
			Subtotal:  subtotal,
		})
	}

	return response, nil
}
