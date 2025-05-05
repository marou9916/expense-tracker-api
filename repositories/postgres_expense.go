package repositories

import (
	"database/sql"
	"errors"

	"github.com/marou9916/expense-tracker-api.git/models"
)

// Implementation of expense interface or "how it works exactly"
type PostgresExpenseRepository struct {
	DB *sql.DB
}

func NewPostgresExpenseRepository(db *sql.DB) *PostgresExpenseRepository {
	return &PostgresExpenseRepository{DB: db}
}

func (r *PostgresExpenseRepository) AddExpense(userUUID string, expense models.Expense) error {
	_, err := r.DB.Exec(
		"INSERT INTO expense(description_expense, amount, date_expense, id_user_uuid) VALUES ($1, $2, $3, $4)",
		expense.Description, expense.Amount, expense.Date, userUUID,
	)
	return err
}

func (r *PostgresExpenseRepository) DeleteExpense(userUUID string, expenseID string) error {
	_, err := r.DB.Exec(
		"DELETE FROM expense WHERE id_expense = $1 AND id_user_uuid = $2", expenseID, userUUID,
	)
	return err
}

func (r *PostgresExpenseRepository) GetExpenses(userUUID string) ([]models.Expense, error) {
	rows, err := r.DB.Query(
		"SELECT * FROM expense WHERE id_user_uuid = $1", userUUID,
	)
	if err != nil {
		return nil, errors.New("failed to get expenses")
	}

	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var expense models.Expense

		if err := rows.Scan(&expense.Description, &expense.Amount, &expense.Date); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)

	}
	return expenses, nil
}
