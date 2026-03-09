package main

import (
	"log"

	"gin-tenant/infrastructure"
	"gin-tenant/middleware"
	"gin-tenant/registry"

	"github.com/gin-gonic/gin"

	_ "gin-tenant/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Multitenant Marketplace API
// @version         1.0
// @description     A Proof of Concept API for multitenant e-commerce.
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token .
func main() {
	cfg := infrastructure.ConfigLoad()
	db := infrastructure.ConnectDB(cfg.DatabaseDSN)
	infrastructure.RunMigrations(db)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	reg := registry.NewRegistry(db)
	reg.RegisterAll(r.Group("/api/v1"))

	log.Printf("Starting server on port %s...", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
