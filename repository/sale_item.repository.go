package repository

import (
	"gorm.io/gorm"
)

type SaleItemRepository interface {
	BaseRepository[SaleItem]
}

type saleItemRepository struct {
	BaseRepository[SaleItem]
	db *gorm.DB
}

func NewSaleItemRepository(db *gorm.DB) SaleItemRepository {
	return &saleItemRepository{
		BaseRepository: NewBaseRepository[SaleItem](db),
		db:             db,
	}
}
