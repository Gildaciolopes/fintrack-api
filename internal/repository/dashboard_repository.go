package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Gildaciolopes/fintrack-backend/internal/models"
	"github.com/google/uuid"
)

type DashboardRepository struct {
	db *sql.DB
}

func NewDashboardRepository(db *sql.DB) *DashboardRepository {
	return &DashboardRepository{db: db}
}

func (r *DashboardRepository) GetStats(userID uuid.UUID, startDate, endDate time.Time) (*models.DashboardStats, error) {
	query := `
		SELECT 
			COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as total_income,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) as total_expenses
		FROM transactions
		WHERE user_id = $1 AND date >= $2 AND date <= $3
	`

	fmt.Printf("[DEBUG GetStats] userID: %s, startDate: %s, endDate: %s\n", userID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	var totalIncome, totalExpenses float64
	err := r.db.QueryRow(query, userID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Scan(&totalIncome, &totalExpenses)
	if err != nil {
		return nil, err
	}

	balance := totalIncome - totalExpenses
	savingsRate := 0.0
	if totalIncome > 0 {
		savingsRate = (balance / totalIncome) * 100
	}

	return &models.DashboardStats{
		TotalIncome:   totalIncome,
		TotalExpenses: totalExpenses,
		Balance:       balance,
		SavingsRate:   savingsRate,
	}, nil
}

func (r *DashboardRepository) GetExpensesByCategory(userID uuid.UUID, startDate, endDate time.Time) ([]models.CategoryExpense, error) {
	query := `
		SELECT 
			COALESCE(c.name, 'Uncategorized') as category,
			SUM(t.amount) as amount,
			COALESCE(c.color, '#6366f1') as color
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.user_id = $1 
			AND t.type = 'expense'
			AND t.date >= $2 
			AND t.date <= $3
		GROUP BY c.name, c.color
		ORDER BY amount DESC
	`

	rows, err := r.db.Query(query, userID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.CategoryExpense
	var totalAmount float64

	for rows.Next() {
		var exp models.CategoryExpense
		if err := rows.Scan(&exp.Category, &exp.Amount, &exp.Color); err != nil {
			return nil, err
		}
		totalAmount += exp.Amount
		expenses = append(expenses, exp)
	}

	// Calculate percentages
	for i := range expenses {
		if totalAmount > 0 {
			expenses[i].Percentage = (expenses[i].Amount / totalAmount) * 100
		}
	}

	return expenses, rows.Err()
}

func (r *DashboardRepository) GetMonthlyData(userID uuid.UUID, months int) ([]models.MonthlyData, error) {
	query := `
		SELECT 
			TO_CHAR(date, 'YYYY-MM') as month,
			COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as income,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) as expenses
		FROM transactions
		WHERE user_id = $1 
			AND date >= DATE_TRUNC('month', CURRENT_DATE) - make_interval(months => $2 - 1)
		GROUP BY TO_CHAR(date, 'YYYY-MM')
		ORDER BY month ASC
	`

	rows, err := r.db.Query(query, userID, months)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var monthlyData []models.MonthlyData
	for rows.Next() {
		var data models.MonthlyData
		if err := rows.Scan(&data.Month, &data.Income, &data.Expenses); err != nil {
			return nil, err
		}
		monthlyData = append(monthlyData, data)
	}

	return monthlyData, rows.Err()
}

func (r *DashboardRepository) GetDailyData(userID uuid.UUID, startDate, endDate time.Time) ([]models.DailyData, error) {
	query := `
		SELECT 
			TO_CHAR(date, 'YYYY-MM-DD') as date_str,
			COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as income,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) as expenses
		FROM transactions
		WHERE user_id = $1 
			AND date >= $2 
			AND date <= $3
		GROUP BY date, TO_CHAR(date, 'YYYY-MM-DD')
		ORDER BY date ASC
	`

	rows, err := r.db.Query(query, userID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dailyData []models.DailyData
	for rows.Next() {
		var data models.DailyData
		if err := rows.Scan(&data.Date, &data.Income, &data.Expenses); err != nil {
			return nil, err
		}
		dailyData = append(dailyData, data)
	}

	return dailyData, rows.Err()
}
