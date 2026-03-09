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

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Takes a product JSON payload and saves it to the database
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      dto.CreateProductInput  true  "Product Data"
// @Success      201      {object}  repository.Product
// @Failure      400      {object}  middleware.AppError
// @Router       /api/v1/products/ [post]
func (h *ProductController) CreateProduct(c *gin.Context) {
	var input dto.CreateProductInput
	tenantID := c.MustGet(middleware.ContextTenantID).(uuid.UUID)

	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(middleware.NewBadRequest("invalid request payload", err))
		return
	}

	createdProduct, err := h.productService.CreateProduct(input, tenantID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductController) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	tenantID := c.MustGet(middleware.ContextTenantID).(uuid.UUID)

	productID, err := uuid.Parse(idParam)

	if err != nil {
		_ = c.Error(middleware.NewBadRequest("invalid UUID format", err))
		return
	}

	product, err := h.productService.GetProduct(productID, tenantID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, product)
}
func (h *ProductController) RegisterRoutes(router *gin.RouterGroup) {
	products := router.Group("/products")
	{
		products.POST("/", middleware.RequireAuth("admin"), h.CreateProduct)
		products.GET("/:id", middleware.RequireAuth("admin", "customer"), h.GetProduct)
	}
}
