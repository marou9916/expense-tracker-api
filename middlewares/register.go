package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marou9916/expense-tracker-api.git/models"
	"github.com/marou9916/expense-tracker-api.git/utils"
)

func FormatValidationHandler(c *gin.Context) {
	var userWantingToRegister models.User

	if err := c.ShouldBindJSON(&userWantingToRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format"})
		c.Abort()
		return
	}

	if !utils.ValidateInputFormat(userWantingToRegister.Name, userWantingToRegister.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name or email"})
		c.Abort()
		return
	}

	if !utils.ValidatePassword(userWantingToRegister.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		c.Abort()
		return
	}

	c.Next()
}
