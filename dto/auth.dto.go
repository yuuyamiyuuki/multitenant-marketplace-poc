package dto

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
