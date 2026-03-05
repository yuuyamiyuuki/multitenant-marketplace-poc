package repository

import (
	"gorm.io/gorm"
)

type ClientRepository interface {
	BaseRepository[*Client]
}

type clientRepository struct {
	BaseRepository[*Client]
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{
		BaseRepository: NewBaseRepository[*Client](db),
		db:             db,
	}
}
