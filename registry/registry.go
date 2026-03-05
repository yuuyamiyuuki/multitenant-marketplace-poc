package registry

import (
	"gorm.io/gorm"

	"gin-tenant/controller"
	"gin-tenant/repository"
	"gin-tenant/service"
)

type Registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) *Registry {
	return &Registry{db: db}
}

func (r *Registry) NewProductController() *controller.ProductController {
	repo := repository.NewProductRepository(r.db)
	svc := service.NewProductService(repo)
	return controller.NewProductController(svc)
}

func (r *Registry) NewUserController() *controller.UserController {
	repo := repository.NewUserRepository(r.db)
	svc := service.NewUserService(repo)
	return controller.NewUserController(svc)
}
