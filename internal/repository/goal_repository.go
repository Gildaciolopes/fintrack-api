package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Gildaciolopes/fintrack-api/internal/models"
	"github.com/google/uuid"
)

type GoalRepository struct {
	db *sql.DB
}

func NewGoalRepository(db *sql.DB) *GoalRepository {
	return &GoalRepository{db: db}
}

func (r *GoalRepository) Create(goal *models.FinancialGoal) error {
	query := `
		INSERT INTO financial_goals (id, user_id, title, target_amount, current_amount, deadline, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	goal.ID = uuid.New()
	goal.CreatedAt = time.Now()
	goal.UpdatedAt = time.Now()
	goal.Status = "active"

	return r.db.QueryRow(
		query,
		goal.ID,
		goal.UserID,
		goal.Title,
		goal.TargetAmount,
		goal.CurrentAmount,
		goal.Deadline,
		goal.Status,
		goal.CreatedAt,
		goal.UpdatedAt,
	).Scan(&goal.ID, &goal.CreatedAt, &goal.UpdatedAt)
}

func (r *GoalRepository) GetByID(id, userID uuid.UUID) (*models.FinancialGoal, error) {
	query := `
		SELECT id, user_id, title, target_amount, current_amount, deadline, status, created_at, updated_at
		FROM financial_goals
		WHERE id = $1 AND user_id = $2
	`

	goal := &models.FinancialGoal{}
	err := r.db.QueryRow(query, id, userID).Scan(
		&goal.ID,
		&goal.UserID,
		&goal.Title,
		&goal.TargetAmount,
		&goal.CurrentAmount,
		&goal.Deadline,
		&goal.Status,
		&goal.CreatedAt,
		&goal.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("goal not found")
	}

	return goal, err
}

func (r *GoalRepository) GetAll(userID uuid.UUID, status string) ([]models.FinancialGoal, error) {
	query := `
		SELECT id, user_id, title, target_amount, current_amount, deadline, status, created_at, updated_at
		FROM financial_goals
		WHERE user_id = $1
	`

	args := []interface{}{userID}

	if status != "" {
		query += " AND status = $2"
		args = append(args, status)
	}

	query += " ORDER BY created_at DESC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []models.FinancialGoal
	for rows.Next() {
		var goal models.FinancialGoal
		if err := rows.Scan(
			&goal.ID,
			&goal.UserID,
			&goal.Title,
			&goal.TargetAmount,
			&goal.CurrentAmount,
			&goal.Deadline,
			&goal.Status,
			&goal.CreatedAt,
			&goal.UpdatedAt,
		); err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}

	return goals, rows.Err()
}

func (r *GoalRepository) Update(id, userID uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	updates["updated_at"] = time.Now()

	var setClauses []string
	var args []interface{}
	argPos := 1

	for field, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, argPos))
		args = append(args, value)
		argPos++
	}

	args = append(args, id, userID)

	query := fmt.Sprintf(
		"UPDATE financial_goals SET %s WHERE id = $%d AND user_id = $%d",
		strings.Join(setClauses, ", "),
		argPos,
		argPos+1,
	)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("goal not found")
	}

	return nil
}

func (r *GoalRepository) Delete(id, userID uuid.UUID) error {
	query := "DELETE FROM financial_goals WHERE id = $1 AND user_id = $2"

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("goal not found")
	}

	return nil
}

func (r *GoalRepository) Contribute(id, userID uuid.UUID, amount float64) error {
	query := `
		UPDATE financial_goals 
		SET current_amount = current_amount + $1, 
		    updated_at = $2,
		    status = CASE 
		        WHEN current_amount + $1 >= target_amount THEN 'completed'
		        ELSE status 
		    END
		WHERE id = $3 AND user_id = $4
	`

	result, err := r.db.Exec(query, amount, time.Now(), id, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("goal not found")
	}

	return nil
}
