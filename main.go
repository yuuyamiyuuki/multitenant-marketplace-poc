package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gin-tenant/controller"
	"gin-tenant/middleware"
	"gin-tenant/repository"
	"gin-tenant/service"
)

func main() {
	// TODO use env vars
	dsn := "host=localhost user=postgres password=postgres dbname=marketplace port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	api := r.Group("/api/v1")
	{
		api.POST("/products", productController.CreateProduct)
		api.GET("/products/:id", productController.GetProduct)
	}

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
