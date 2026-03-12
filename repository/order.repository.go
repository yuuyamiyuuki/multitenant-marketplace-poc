package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	BaseRepository[Order]
	FindByClientIDAndTenant(clientID uuid.UUID, tenantID uuid.UUID) ([]*Order, error)
}

type orderRepository struct {
	BaseRepository[Order]
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		BaseRepository: NewBaseRepository[Order](db),
		db:             db,
	}
}

func (repo *orderRepository) FindByClientIDAndTenant(clientID uuid.UUID, tenantID uuid.UUID) ([]*Order, error) {
	var orders []*Order
	result := repo.db.Preload("Items.Product").Where("client_id = ? AND tenant_id = ?", clientID, tenantID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
