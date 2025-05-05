package repositories

import "github.com/marou9916/expense-tracker-api.git/models"

//	Expense interface or "what we can do with the expenses"
type ExpenseRepository interface {
	AddExpense(userUUID string, expense models.Expense)
	DeleteExpense(userUUID string, expense models.Expense)
	GetExpenses(userUUID string) ([]models.Expense, error)
}

