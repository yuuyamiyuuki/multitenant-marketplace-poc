package middleware

import (
	"fmt"
	"gin-tenant/dto"
	"gin-tenant/infrastructure"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	ContextTenantID = "tenant_id"
	ContextUserID   = "user_id"
	ContextUserRole = "user_role"
)

func RequireAuth(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization header"})
			return
		}

		claims, err := validateToken(tokenString)
		if err != nil || claims.Subject == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token or missing identifier"})
			return
		}

		if !isAuthorized(claims.Role, allowedRoles) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		tenantUUID, err := uuid.Parse(claims.TenantID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "corrupt tenant ID in token"})
			return
		}

		userUUID, err := uuid.Parse(claims.Subject)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "corrupt user ID in token"})
			return
		}

		c.Set(ContextTenantID, tenantUUID)
		c.Set(ContextUserID, userUUID)
		c.Set(ContextUserRole, claims.Role)

		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(authHeader, "Bearer ")
}

func validateToken(tokenString string) (*dto.CustomClaims, error) {
	secret := []byte(infrastructure.App.JWTSecret)

	token, err := jwt.ParseWithClaims(tokenString, &dto.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*dto.CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func isAuthorized(userRole string, allowedRoles []string) bool {
	if len(allowedRoles) == 0 {
		return true
	}
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
