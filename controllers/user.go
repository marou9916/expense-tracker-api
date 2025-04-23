package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marou9916/expense-tracker-api.git/database"
	"github.com/marou9916/expense-tracker-api.git/models"
)

func GetCurrentUser(c *gin.Context) {
	userUUID := c.MustGet("userUUID").(string)

	rows, err := database.DB.Query(
		"SELECT description_expense, amount, date_expense FROM expense WHERE id_user_uuid = $1", userUUID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible to continue the request"})
		return
	}

	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var e models.Expense
		if err := rows.Scan(&e.Description, &e.Amount, &e.Date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to show your expenses"})
			return
		}

		expenses = append(expenses, e)
	}

	c.JSON(http.StatusOK, gin.H{"Expenses": expenses})
}
