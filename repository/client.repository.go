package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientRepository interface {
	BaseRepository[Client]

	FindByUserIDAndTenant(userID uuid.UUID, tenantID uuid.UUID) (*Client, error)
}

type clientRepository struct {
	BaseRepository[Client]
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{
		BaseRepository: NewBaseRepository[Client](db),
		db:             db,
	}
}

func (repo *clientRepository) FindByUserIDAndTenant(userID uuid.UUID, tenantID uuid.UUID) (*Client, error) {
	var client Client
	result := repo.db.Where("user_id = ? AND tenant_id = ?", userID, tenantID).First(&client)
	if result.Error != nil {
		return nil, result.Error
	}
	return &client, nil
}
