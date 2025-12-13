package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Gildaciolopes/fintrack-backend/internal/models"
	"github.com/google/uuid"
)

type BudgetRepository struct {
	db *sql.DB
}

func NewBudgetRepository(db *sql.DB) *BudgetRepository {
	return &BudgetRepository{db: db}
}

func (r *BudgetRepository) Create(budget *models.Budget) error {
	query := `
		INSERT INTO budgets (id, user_id, category_id, amount, month, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	budget.ID = uuid.New()
	budget.CreatedAt = time.Now()

	return r.db.QueryRow(
		query,
		budget.ID,
		budget.UserID,
		budget.CategoryID,
		budget.Amount,
		budget.Month,
		budget.CreatedAt,
	).Scan(&budget.ID, &budget.CreatedAt)
}

func (r *BudgetRepository) GetByID(id, userID uuid.UUID) (*models.Budget, error) {
	query := `
		SELECT 
			b.id, b.user_id, b.category_id, b.amount, b.month, b.created_at,
			c.id, c.user_id, c.name, c.type, c.color, c.icon, c.created_at
		FROM budgets b
		LEFT JOIN categories c ON b.category_id = c.id
		WHERE b.id = $1 AND b.user_id = $2
	`

	budget := &models.Budget{}
	var category models.Category

	err := r.db.QueryRow(query, id, userID).Scan(
		&budget.ID,
		&budget.UserID,
		&budget.CategoryID,
		&budget.Amount,
		&budget.Month,
		&budget.CreatedAt,
		&category.ID,
		&category.UserID,
		&category.Name,
		&category.Type,
		&category.Color,
		&category.Icon,
		&category.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("budget not found")
	}
	if err != nil {
		return nil, err
	}

	budget.Category = &category
	return budget, nil
}

func (r *BudgetRepository) GetAll(userID uuid.UUID, month *time.Time) ([]models.Budget, error) {
	query := `
		SELECT 
			b.id, b.user_id, b.category_id, b.amount, b.month, b.created_at,
			c.id, c.user_id, c.name, c.type, c.color, c.icon, c.created_at
		FROM budgets b
		LEFT JOIN categories c ON b.category_id = c.id
		WHERE b.user_id = $1
	`

	args := []interface{}{userID}

	if month != nil {
		query += " AND DATE_TRUNC('month', b.month) = DATE_TRUNC('month', $2::date)"
		args = append(args, *month)
	}

	query += " ORDER BY b.month DESC, c.name ASC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var budgets []models.Budget
	for rows.Next() {
		var budget models.Budget
		var category models.Category

		if err := rows.Scan(
			&budget.ID,
			&budget.UserID,
			&budget.CategoryID,
			&budget.Amount,
			&budget.Month,
			&budget.CreatedAt,
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.Type,
			&category.Color,
			&category.Icon,
			&category.CreatedAt,
		); err != nil {
			return nil, err
		}

		budget.Category = &category
		budgets = append(budgets, budget)
	}

	return budgets, rows.Err()
}

func (r *BudgetRepository) GetBudgetsWithSpent(userID uuid.UUID, month time.Time) ([]models.BudgetWithSpent, error) {
	query := `
		SELECT 
			b.id, b.user_id, b.category_id, b.amount, b.month, b.created_at,
			c.id, c.user_id, c.name, c.type, c.color, c.icon, c.created_at,
			COALESCE(SUM(t.amount), 0) as spent
		FROM budgets b
		LEFT JOIN categories c ON b.category_id = c.id
		LEFT JOIN transactions t ON t.category_id = b.category_id 
			AND t.user_id = b.user_id 
			AND t.type = 'expense'
			AND DATE_TRUNC('month', t.date) = DATE_TRUNC('month', b.month)
		WHERE b.user_id = $1 
			AND DATE_TRUNC('month', b.month) = DATE_TRUNC('month', $2::date)
		GROUP BY b.id, b.user_id, b.category_id, b.amount, b.month, b.created_at,
				 c.id, c.user_id, c.name, c.type, c.color, c.icon, c.created_at
		ORDER BY c.name ASC
	`

	rows, err := r.db.Query(query, userID, month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var budgetsWithSpent []models.BudgetWithSpent
	for rows.Next() {
		var bws models.BudgetWithSpent
		var category models.Category

		if err := rows.Scan(
			&bws.ID,
			&bws.UserID,
			&bws.CategoryID,
			&bws.Amount,
			&bws.Month,
			&bws.CreatedAt,
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.Type,
			&category.Color,
			&category.Icon,
			&category.CreatedAt,
			&bws.Spent,
		); err != nil {
			return nil, err
		}

		bws.Category = &category
		bws.Remaining = bws.Amount - bws.Spent
		if bws.Amount > 0 {
			bws.Percentage = (bws.Spent / bws.Amount) * 100
		}

		budgetsWithSpent = append(budgetsWithSpent, bws)
	}

	return budgetsWithSpent, rows.Err()
}

func (r *BudgetRepository) Update(id, userID uuid.UUID, amount float64, month time.Time) error {
	query := `
		UPDATE budgets 
		SET amount = $1, month = $2
		WHERE id = $3 AND user_id = $4
	`

	result, err := r.db.Exec(query, amount, month, id, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("budget not found")
	}

	return nil
}

func (r *BudgetRepository) Delete(id, userID uuid.UUID) error {
	query := "DELETE FROM budgets WHERE id = $1 AND user_id = $2"

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("budget not found")
	}

	return nil
}
