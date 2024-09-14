package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh-solanki21/golang-gin-crud-api/utils"
)

func AuthorizeMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := getClaimsFromContext(c)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		}

		if !isAuthorized(claims, roles) {
			utils.RespondWithError(c, http.StatusForbidden, "Access denied", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func getClaimsFromContext(c *gin.Context) (*utils.Claims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "Unauthorized", nil)
	}

	userClaims, ok := claims.(*utils.Claims)
	if !ok {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Invalid claims", nil)
	}

	return userClaims, nil
}

func isAuthorized(claims *utils.Claims, roles []string) bool {
	if len(roles) == 0 {
		return true
	}

	for _, role := range roles {
		if claims.Role == role {
			return true
		}
	}

	return false
}
