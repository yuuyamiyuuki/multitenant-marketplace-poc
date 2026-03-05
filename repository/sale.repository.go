package repository

import (
	"gorm.io/gorm"
)

type SaleRepository interface {
	BaseRepository[Sale]
}

type saleRepository struct {
	BaseRepository[Sale]
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) SaleRepository {
	return &saleRepository{
		BaseRepository: NewBaseRepository[Sale](db),
		db:             db,
	}
}
