package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	FindByID(id uuid.UUID) (*T, error)
	FindByIDAndTenant(id uuid.UUID, tenantID uuid.UUID) (*T, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

func (repo *baseRepository[T]) Create(entity *T) error {
	return repo.db.Create(entity).Error
}

func (repo *baseRepository[T]) Update(entity *T) error {
	return repo.db.Save(entity).Error
}

func (repo *baseRepository[T]) Delete(entity *T) error {
	return repo.db.Delete(entity).Error
}

func (repo *baseRepository[T]) FindByID(id uuid.UUID) (*T, error) {
	var entity T

	result := repo.db.First(&entity, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, result.Error
	}

	return &entity, nil
}

func (repo *baseRepository[T]) FindByIDAndTenant(id uuid.UUID, tenantID uuid.UUID) (*T, error) {
	var entity T

	result := repo.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&entity)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, result.Error
	}

	return &entity, nil
}
