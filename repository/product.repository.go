package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	BaseRepository[Product]

	FindByIDAndTenant(id uuid.UUID, tenantID uuid.UUID) (*Product, error)
}

type productRepository struct {
	BaseRepository[Product]
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		BaseRepository: NewBaseRepository[Product](db),
		db:             db,
	}
}
