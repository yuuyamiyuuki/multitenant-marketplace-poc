package registry

import (
	"gin-tenant/controller"
	"gin-tenant/repository"
	"gin-tenant/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller interface {
	RegisterRoutes(router *gin.RouterGroup)
}

type Registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) *Registry {
	return &Registry{db: db}
}

func (r *Registry) RegisterAll(router *gin.RouterGroup) {
	controllers := []Controller{
		r.NewUserController(),
		r.NewProductController(),
	}

	for _, ctrl := range controllers {
		ctrl.RegisterRoutes(router)
	}
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
