package service

import (
	"github.com/google/uuid"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/repository"
)

type ProductService interface {
	CreateProduct(input dto.CreateProductInput, tenantID uuid.UUID) (*repository.Product, error)
	GetProduct(id uuid.UUID, tenantID uuid.UUID) (*repository.Product, error)
	GetProductsByIDs(ids []uuid.UUID, tenantID uuid.UUID) ([]*repository.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: repo,
	}
}

func (s *productService) CreateProduct(input dto.CreateProductInput, tenantID uuid.UUID) (*repository.Product, error) {
	newProduct := &repository.Product{
		TenantID: tenantID,
		Name:     input.Name,
		Price:    input.Price,
		Stock:    input.Stock,
	}

	err := s.productRepo.Create(newProduct)
	if err != nil {
		return nil, middleware.NewInternal("failed to create product in the database", err)
	}

	return newProduct, nil
}

func (s *productService) GetProduct(id uuid.UUID, tenantID uuid.UUID) (*repository.Product, error) {

	product, err := s.productRepo.FindByIDAndTenant(id, tenantID)

	if err != nil {
		if err.Error() == "record not found" {
			return nil, middleware.NewNotFound("the requested product does not exist", err)
		}
		return nil, middleware.NewInternal("failed to fetch product from database", err)
	}

	return product, nil
}

func (s *productService) GetProductsByIDs(ids []uuid.UUID, tenantID uuid.UUID) ([]*repository.Product, error) {
	products, err := s.productRepo.FindByIDListAndTenant(ids, tenantID)
	if err != nil {
		return nil, middleware.NewInternal("failed to fetch products from database", err)
	}
	return products, nil
}
