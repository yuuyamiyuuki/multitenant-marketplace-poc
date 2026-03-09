package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Creates a new tenant user. Requires admin privileges.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user  body      dto.CreateUserInput  true  "User Data"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  middleware.AppError
// @Failure      401   {object}  middleware.AppError
// @Failure      403   {object}  middleware.AppError
// @Failure      500   {object}  middleware.AppError
// @Router       /api/v1/users/ [post]
func (h *UserController) CreateUser(c *gin.Context) {
	tenantID := c.MustGet(middleware.ContextTenantID).(uuid.UUID)

	var input dto.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(middleware.NewBadRequest("invalid request payload", err))
		return
	}

	createdUser, err := h.userService.CreateTenantUser(input, tenantID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    createdUser,
	})
}

// UpdatePassword godoc
// @Summary      Update current user password
// @Description  Updates the password for the authenticated user. User ID is extracted from JWT.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        payload  body      dto.UpdatePasswordInput  true  "Password Data"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  middleware.AppError
// @Failure      401      {object}  middleware.AppError
// @Failure      500      {object}  middleware.AppError
// @Router       /api/v1/users/password [patch]
func (h *UserController) UpdatePassword(c *gin.Context) {
	tenantID := c.MustGet(middleware.ContextTenantID).(uuid.UUID)
	userID := c.MustGet(middleware.ContextUserID).(uuid.UUID)

	var input dto.UpdatePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(middleware.NewBadRequest("invalid request payload", err))
		return
	}

	err := h.userService.UpdatePassword(userID, tenantID, input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// Login godoc
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        credentials  body      dto.LoginInput  true  "Login Credentials"
// @Success      200          {object}  dto.LoginResponse
// @Failure      400          {object}  middleware.AppError
// @Failure      401          {object}  middleware.AppError
// @Failure      500          {object}  middleware.AppError
// @Router       /api/v1/users/login [post]
func (h *UserController) Login(c *gin.Context) {
	var input dto.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(middleware.NewBadRequest("invalid request payload", err))
		return
	}

	response, err := h.userService.Login(input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserController) RegisterRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.POST("/login", h.Login)
		users.POST("/", middleware.RequireAuth("admin"), h.CreateUser)
		users.PATCH("/password", middleware.RequireAuth("admin", "customer"), h.UpdatePassword)
	}
}
