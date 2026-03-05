package repository

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
	BaseRepository[Product]
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
