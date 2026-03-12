package repository

import (
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	BaseRepository[OrderItem]
}

type orderItemRepository struct {
	BaseRepository[OrderItem]
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{
		BaseRepository: NewBaseRepository[OrderItem](db),
		db:             db,
	}
}
