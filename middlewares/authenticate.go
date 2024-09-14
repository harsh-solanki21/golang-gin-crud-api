package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh-solanki21/golang-gin-crud-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			handleTokenError(c, err)
			return
		}

		claims, err := utils.ValidateAccessToken(accessToken)
		if err != nil {
			// Access token is invalid or expired, try to refresh
			refreshToken, err := c.Cookie("refresh_token")
			if err != nil {
				handleTokenError(c, err)
				return
			}

			newAccessToken, err := utils.RefreshAccessToken(refreshToken)
			if err != nil {
				handleTokenError(c, err)
				return
			}

			// Set the new access token as a cookie
			c.SetCookie("access_token", newAccessToken, 900, "/", "", false, true)

			claims, err = utils.ValidateAccessToken(newAccessToken)
			if err != nil {
				handleTokenError(c, err)
				return
			}
		}

		// Set the entire claims object in the context
		c.Set("claims", claims)
		log.Println("User ID:", claims.UserID, "Role:", claims.Role)

		c.Next()
	}
}

func handleTokenError(c *gin.Context, err error) {
	utils.RespondWithError(c, http.StatusUnauthorized, "Authentication failed", err)
	c.Abort()
}
