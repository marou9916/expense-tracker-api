package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marou9916/expense-tracker-api.git/database"
	"github.com/marou9916/expense-tracker-api.git/models"
	"github.com/marou9916/expense-tracker-api.git/utils"
	"golang.org/x/crypto/bcrypt"
)

// Handlers that handle the user registration/login process.
func RegisterHandler(c *gin.Context) {

}

func LoginHandler(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var userUUID string
	var hashedPassword string

	err := database.DB.QueryRow(
		"SELECT id_user_uuid password_user FROM users WHERE email_user = $1", input.Email,
	).Scan(&userUUID, &hashedPassword)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user with that email found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	tokenJWT, err := utils.GenerateJWT(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenJWT})

}
