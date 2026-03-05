package service

import (
	"github.com/google/uuid"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/repository"
)

type ProductService interface {
	CreateProduct(input dto.CreateProductInput) (*repository.Product, error)
	GetProduct(id uuid.UUID) (*repository.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: repo,
	}
}

func (s *productService) CreateProduct(input dto.CreateProductInput) (*repository.Product, error) {
	newProduct := &repository.Product{
		Name:  input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	err := s.productRepo.Create(newProduct)
	if err != nil {
		return nil, middleware.NewInternal("failed to create product in the database", err)
	}

	return newProduct, nil
}

func (s *productService) GetProduct(id uuid.UUID) (*repository.Product, error) {
	product, err := s.productRepo.FindByID(id)

	if err != nil {
		if err.Error() == "record not found" {
			return nil, middleware.NewNotFound("the requested product does not exist", err)
		}
		return nil, middleware.NewInternal("failed to fetch product from database", err)
	}

	return product, nil
}
