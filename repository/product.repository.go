package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	BaseRepository[Product]

	FindByIDAndTenant(id uuid.UUID, tenantID uuid.UUID) (*Product, error)
	FindByIDListAndTenant(idList []uuid.UUID, tenantID uuid.UUID) ([]*Product, error)
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

func (repo *productRepository) FindByIDListAndTenant(idList []uuid.UUID, tenantID uuid.UUID) ([]*Product, error) {
	var products []*Product
	result := repo.db.Where("id IN (?) AND tenant_id = ?", idList, tenantID).Find(&products)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, result.Error
	}

	return products, nil
}
