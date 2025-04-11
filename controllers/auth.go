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
	var registration models.RegisterInput

	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !utils.ValidateInputFormat(registration.Name, registration.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or email format"})
		return
	}

	if !utils.ValidatePassword(registration.Password) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password must be at least 8 characters long and contain both letters and digits"})
		return
	}

	var existingEmail string

	err := database.DB.QueryRow(
		"SELECT email_user FROM users WHERE email_user = $1", registration.Email,
	).Scan(&existingEmail)

	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if existingEmail != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already used"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save password"})
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO users(name_user, email_user, password_user) VALUES($1, $2, $3)",
		registration.Name, registration.Email, string(hashedPassword),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user succesfully registered"})
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
