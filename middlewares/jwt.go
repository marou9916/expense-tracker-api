package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/marou9916/expense-tracker-api.git/utils"
)

// JWTRequired is a middleware that ensures a valid JWT token is present.
// It verifies the user's identity by checking the token in the Authorization header.
func JWTRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized to continue the request"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(token, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
			c.Abort()
			return
		}

		tokenValueJWT := tokenParts[1]

		if !utils.VerifyJWT(tokenValueJWT) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		

	}
}
