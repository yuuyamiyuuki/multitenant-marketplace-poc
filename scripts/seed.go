//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/google/uuid"

	"gin-tenant/infrastructure"
	"gin-tenant/repository"
	"gin-tenant/utils"
)

func main() {
	fmt.Println("Starting postgres")
	exec.Command("docker-compose", "up", "-d").Run()

	fmt.Println("Waiting for postgres")
	time.Sleep(3 * time.Second)

	cfg := infrastructure.ConfigLoad()
	db := infrastructure.ConnectDB(cfg.DatabaseDSN)
	infrastructure.RunMigrations(db)

	var count int64
	db.Model(&repository.User{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded. Seeding skipped.")
		return
	}

	adminUser := os.Getenv("ADMIN_USER")
	adminPass := os.Getenv("ADMIN_PASSWORD")

	if adminUser == "" || adminPass == "" {
		log.Fatal("ADMIN_USER and ADMIN_PASSWORD must be set in the environment to run the seeder")
	}

	hashedPassword, err := utils.HashPassword(adminPass)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	tenant := repository.Tenant{
		Name:   "admin",
		APIKey: uuid.New(),
	}
	if err := db.Create(&tenant).Error; err != nil {
		log.Fatalf("Failed to create root tenant: %v", err)
	}

	admin := repository.User{
		TenantID: tenant.ID,
		Username: adminUser,
		Role:     "admin",
		Password: hashedPassword,
	}
	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	log.Println("Initial Admin User seeded successfully")
	log.Printf("Tenant ID: %s", tenant.ID.String())
	log.Printf("Username: %s", adminUser)
}
