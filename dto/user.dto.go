package dto

import "time"

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdatePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
