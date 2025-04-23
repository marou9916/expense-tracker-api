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
	userInput, exists := c.Get("userInput")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected middleware failure"})
		return
	}

	input := userInput.(models.RegisterInput)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password),  bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Unexpected error during registration process"})
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO users(name_user, email_user, password_user) VALUES ($1, $2, $3)", input.Name, input.Email, string(hashedPassword),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register the request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully registered"})
}

func LoginHandler(c *gin.Context) {
	userInput, exists := c.Get("userInput")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected middleware failure"})
		return
	}

	input := userInput.(models.LoginInput)

	var userUUID string
	var hashedPassword string

	err := database.DB.QueryRow(
		"SELECT id_user_uuid, password_user FROM users WHERE email_user = $1", input.Email,
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to complete login process"})
		return
	}

	c.Header("Authorization", "Bearer "+tokenJWT)

	c.JSON(http.StatusOK, gin.H{"message": "User successfully logged"})

}
