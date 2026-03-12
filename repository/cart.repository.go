package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	BaseRepository[Cart]

	FindByIDAndTenant(id uuid.UUID, tenantID uuid.UUID) (*Cart, error)
}

type cartRepository struct {
	BaseRepository[Cart]
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		BaseRepository: NewBaseRepository[Cart](db),
		db:             db,
	}
}
