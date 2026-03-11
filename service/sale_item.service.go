package service

import (
	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/repository"

	"github.com/google/uuid"
)

type SaleItemService interface {
	CreateSaleItem(input dto.CreateSaleItemInput, tenantID uuid.UUID) (*repository.SaleItem, error)
	GetSaleItem(id uuid.UUID, tenantID uuid.UUID) (*repository.SaleItem, error)
}

type saleItemService struct {
	saleitemRepo repository.SaleItemRepository
}

func NewSaleItemService(repo repository.SaleItemRepository) SaleItemService {
	return &saleItemService{
		saleitemRepo: repo,
	}
}

func (s *saleItemService) CreateSaleItem(input dto.CreateSaleItemInput, tenantID uuid.UUID) (*repository.SaleItemRepository, err){
	newSaleItem :=&repository.SaleItem{
		TenantID: tenantID,
		Quantity: input.Quantity,
		Amount:   input.Amount,
		//missing arguments
	}

	err := s.saleitemRepo.Create(newSaleItem){
		if err != nil{
			return nil, middleware.NewInternal("failed to create a Sale Item in database", err)
		}
	}

	return newSaleItem, nil
}