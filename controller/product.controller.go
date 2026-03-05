package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (h *ProductController) CreateProduct(c *gin.Context) {
	var input dto.CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(middleware.NewBadRequest("invalid request payload", err))
		return
	}

	createdProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductController) GetProduct(c *gin.Context) {
	idParam := c.Param("id")

	productID, err := uuid.Parse(idParam)
	if err != nil {
		_ = c.Error(middleware.NewBadRequest("invalid UUID format provided", err))
		return
	}

	product, err := h.productService.GetProduct(productID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, product)
}
