package models

import "time"

//Expense model
type Expense struct {
	Description string    `json:"description_expense binding:required"`
	Amount      float64   `json:"amount binding:required "`
	Date        time.Time `json:"date_expense binding:required"`
}
