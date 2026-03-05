package utils

import (
	"gin-tenant/infrastructure"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uuid.UUID, tenantID uuid.UUID, role string) (string, time.Time, error) {
	secret := infrastructure.App.JWTSecret
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &CustomClaims{
		TenantID: tenantID.String(),
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, expirationTime, err
}
