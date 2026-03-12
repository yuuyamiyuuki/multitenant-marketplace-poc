package infrastructure

import (
	"log"

	"gorm.io/gorm"

	"gin-tenant/repository"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&repository.Tenant{},
		&repository.User{},
		&repository.Client{},
		&repository.Product{},
		&repository.Order{},
		&repository.OrderItem{},
		&repository.Cart{},
		&repository.CartItem{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migrations completed successfully")
}
