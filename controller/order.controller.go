package controller

import (
	"net/http"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(svc service.OrderService) *OrderController {
	return &OrderController{orderService: svc}
}

func (ctrl *OrderController) RegisterRoutes(router *gin.RouterGroup) {
	checkoutGroup := router.Group("/checkout").Use(middleware.RequireAuth("customer"))
	{
		checkoutGroup.POST("", ctrl.Checkout)
	}

	ordersGroup := router.Group("/orders").Use(middleware.RequireAuth("customer"))
	{
		ordersGroup.GET("", ctrl.ListMyOrders)
	}

	// Webhook should be public or use a different auth structure.
	webhookGroup := router.Group("/webhook")
	{
		webhookGroup.POST("/payment", ctrl.WebhookPayment)
	}
}

// Checkout godoc
// @Summary      Checkout User Cart
// @Description  Processes the user's cart into a final order, deducts product stock, and empties the cart.
// @Tags         Checkout
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      201 {object} repository.Order
// @Router       /checkout [post]
func (ctrl *OrderController) Checkout(c *gin.Context) {
	tenantIDRaw, _ := c.Get("tenant_id")
	userIDRaw, _ := c.Get("user_id")

	tenantID, _ := uuid.Parse(tenantIDRaw.(string))
	userID, _ := uuid.Parse(userIDRaw.(string))

	var input dto.CheckoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload or unaccepted payment method"})
		return
	}

	order, err := ctrl.orderService.Checkout(input, tenantID, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (ctrl *OrderController) ListMyOrders(c *gin.Context) {
	tenantIDRaw, _ := c.Get("tenant_id")
	userIDRaw, _ := c.Get("user_id")

	tenantID, _ := uuid.Parse(tenantIDRaw.(string))
	userID, _ := uuid.Parse(userIDRaw.(string))

	orders, err := ctrl.orderService.ListMyOrders(tenantID, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (ctrl *OrderController) WebhookPayment(c *gin.Context) {
	// A mock representation of a payment provider hitting our webhook
	var payload struct {
		OrderID string `json:"order_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid webhook payload"})
		return
	}

	orderUUID, err := uuid.Parse(payload.OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order_id format"})
		return
	}

	if err := ctrl.orderService.ConcludeOrder(orderUUID); err != nil {
		// Log the error but typical webhooks expect an ACK (200 OK) if it was received properly to stop retries.
		// Sending 400 for bad data or states to help debug.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment concluded and registered successfully"})
}
