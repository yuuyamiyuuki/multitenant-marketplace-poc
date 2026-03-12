package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	BaseRepository[Cart]

	FindByIDAndTenant(id uuid.UUID, tenantID uuid.UUID) (*Cart, error)
	FindByUserIDAndTenant(userID uuid.UUID, tenantID uuid.UUID) (*Cart, error)
	ClearItems(cartID uuid.UUID, tenantID uuid.UUID) error
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

func (repo *cartRepository) FindByUserIDAndTenant(userID uuid.UUID, tenantID uuid.UUID) (*Cart, error) {
	var cart Cart
	result := repo.db.Preload("Items").Where("user_id = ? AND tenant_id = ?", userID, tenantID).First(&cart)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cart, nil
}

func (repo *cartRepository) ClearItems(cartID uuid.UUID, tenantID uuid.UUID) error {
	return repo.db.Where("cart_id = ?", cartID).Delete(&CartItem{}).Error
}
