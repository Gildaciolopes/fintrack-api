package models

type DashboardStats struct {
	TotalIncome   float64 `json:"totalIncome"`
	TotalExpenses float64 `json:"totalExpenses"`
	Balance       float64 `json:"balance"`
	SavingsRate   float64 `json:"savingsRate"`
}

type CategoryExpense struct {
	Category   string  `json:"category" db:"category"`
	Amount     float64 `json:"amount" db:"amount"`
	Color      string  `json:"color" db:"color"`
	Percentage float64 `json:"percentage"`
}

type MonthlyData struct {
	Month    string  `json:"month" db:"month"`
	Income   float64 `json:"income" db:"income"`
	Expenses float64 `json:"expenses" db:"expenses"`
}

type DailyData struct {
	Date     string  `json:"date" db:"date"`
	Income   float64 `json:"income" db:"income"`
	Expenses float64 `json:"expenses" db:"expenses"`
}
