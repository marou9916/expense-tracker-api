// middlewares/validate_input.go

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marou9916/expense-tracker-api.git/models"
	"github.com/marou9916/expense-tracker-api.git/utils"
)

func ValidateInputFormat(inputType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch inputType {
		case "register":
			var input models.RegisterInput

			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format"})
				c.Abort()
				return
			}

			if !utils.ValidateInputFormat(input.Name, input.Email) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name or email"})
				c.Abort()
				return
			}

			if !utils.ValidatePassword(input.Password) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
				c.Abort()
				return
			}

			// Réinjecter l’input dans le contexte si besoin
			c.Set("userInput", input)

		case "login":
			var input models.LoginInput

			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format"})
				c.Abort()
				return
			}

			if !utils.ValidateInputFormat("", input.Email) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
				c.Abort()
				return
			}

			if !utils.ValidatePassword(input.Password) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
				c.Abort()
				return
			}

			c.Set("userInput", input)
		}

		c.Next()
	}
}
