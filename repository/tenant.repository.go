package repository

import (
	"gorm.io/gorm"
)

type TenantRepository interface {
	BaseRepository[Tenant]
}

type tenantRepository struct {
	BaseRepository[Tenant]
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{
		BaseRepository: NewBaseRepository[Tenant](db),
		db:             db,
	}
}
