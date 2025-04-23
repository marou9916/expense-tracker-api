package models

import (
	"errors"

	"github.com/marou9916/expense-tracker-api.git/database"
)

type User struct {
	Name     string `json:"name_user" binding:"required, email"`
	Email    string `json:"email_user" binding:"required, email"`
	Password string `json:"password_user" binding:"required"`
	Expenses []Expense
}

func (u *User) AddExpense(expense Expense) error {
	_, err := database.DB.Exec(
		"INSERT INTO expense(description_expense, amount, date_expense) VALUES ($1, $2, $3)", expense.Description, expense.Amount, expense.Date,
	)
	return err
}

func (u *User) DeleteExpense(expense Expense) error {
	_, err := database.DB.Exec(
		"DELETE * FROM expense WHERE description_expense = $1", expense.Description,
	)

	return err
}

func (u *User) GetExpense(uuid string) ([]Expense, error) {
	var expenses []Expense

	rows, err := database.DB.Query(
		"SELECT description_expense, amount, date_expense FROM Expenses WHERE  id_user_uuid = $1 LIMIT 20", uuid,
	)

	if err != nil {
		return nil, errors.New("failed to get expenses")
	}

	defer rows.Close()

	for rows.Next() {
		var expense Expense

		if err := rows.Scan(&expense.Description, &expense.Amount, &expense.Date); err != nil {
			return nil, errors.New("failed to show your expenses")
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
