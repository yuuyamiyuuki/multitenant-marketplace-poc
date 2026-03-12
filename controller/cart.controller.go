package controller

import (
	"log"
	"net/http"

	"gin-tenant/dto"
	"gin-tenant/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartController struct {
	cartService service.CartService
}

func NewCartController(svc service.CartService) *CartController {
	return &CartController{cartService: svc}
}

func (ctrl *CartController) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/cart")
	{
		group.GET("", ctrl.GetCart)
		group.POST("", ctrl.UpsertCart)
	}
}

// GetCart godoc
// @Summary      Get the user cart
// @Description  Retrive cart and current product items and prices for the current user token
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} dto.CartResponse
// @Router       /cart [get]
func (ctrl *CartController) GetCart(c *gin.Context) {
	tenantIDRaw, _ := c.Get("tenant_id")
	userIDRaw, _ := c.Get("user_id")

	tenantID, _ := uuid.Parse(tenantIDRaw.(string))
	userID, _ := uuid.Parse(userIDRaw.(string))

	cart, err := ctrl.cartService.GetCart(tenantID, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, cart)
}

// UpsertCart godoc
// @Summary      Upsert user cart items
// @Description  Adds or overwrites the current user's cart contents with new items and quantities
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input body dto.UpsertCartInput true "Cart Items Payload"
// @Success      200
// @Router       /cart [post]
func (ctrl *CartController) UpsertCart(c *gin.Context) {
	tenantIDRaw, _ := c.Get("tenant_id")
	userIDRaw, _ := c.Get("user_id")

	tenantID, _ := uuid.Parse(tenantIDRaw.(string))
	userID, _ := uuid.Parse(userIDRaw.(string))

	var input dto.UpsertCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("invalid request payload %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	_, err := ctrl.cartService.UpsertCart(input, tenantID, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cart updated successfully"})
}
