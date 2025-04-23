package models

type Expense struct {
	Description string  `json:"description_expense binding:required"`
	Amount      float64 `json:"amount binding:required "`
	Date        string  `json:"date_expense binding:required"`
}
