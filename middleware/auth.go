package middleware

import (
	"gin-tenant/infrastructure"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("my-super-secret-key")

func RequireAuth(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := infrastructure.App.JWTSecret
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		if claims.Subject == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token missing user identifier"})
			return
		}

		if len(allowedRoles) > 0 {
			hasRole := false
			for _, role := range allowedRoles {
				if claims.Role == role {
					hasRole = true
					break
				}
			}
			if !hasRole {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to access this resource"})
				return
			}
		}

		c.Set("tenant_id", claims.TenantID)
		c.Set("user_id", claims.Subject)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}
