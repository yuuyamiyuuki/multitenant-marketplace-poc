package repository

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[User]
	FindByUsername(username string) (*User, error)
}

type userRepository struct {
	BaseRepository[User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[User](db),
		db:             db,
	}
}

func (repo *userRepository) FindByUsername(username string) (*User, error) {
	var user User

	result := repo.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}
