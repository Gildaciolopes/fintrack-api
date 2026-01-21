package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Gildaciolopes/fintrack-api/internal/models"
	"github.com/google/uuid"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
	query := `
		INSERT INTO categories (id, user_id, name, type, color, icon, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	category.ID = uuid.New()
	category.CreatedAt = time.Now()

	return r.db.QueryRow(
		query,
		category.ID,
		category.UserID,
		category.Name,
		category.Type,
		category.Color,
		category.Icon,
		category.CreatedAt,
	).Scan(&category.ID, &category.CreatedAt)
}

func (r *CategoryRepository) GetByID(id, userID uuid.UUID) (*models.Category, error) {
	query := `
		SELECT id, user_id, name, type, color, icon, created_at
		FROM categories
		WHERE id = $1 AND user_id = $2
	`

	category := &models.Category{}
	err := r.db.QueryRow(query, id, userID).Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
		&category.Type,
		&category.Color,
		&category.Icon,
		&category.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("category not found")
	}

	return category, err
}

func (r *CategoryRepository) GetAll(userID uuid.UUID, categoryType string) ([]models.Category, error) {
	query := `
		SELECT id, user_id, name, type, color, icon, created_at
		FROM categories
		WHERE user_id = $1
	`

	args := []interface{}{userID}

	if categoryType != "" {
		query += " AND type = $2"
		args = append(args, categoryType)
	}

	query += " ORDER BY name ASC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(
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
		categories = append(categories, category)
	}

	return categories, rows.Err()
}

func (r *CategoryRepository) Update(id, userID uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

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
		"UPDATE categories SET %s WHERE id = $%d AND user_id = $%d",
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
		return fmt.Errorf("category not found")
	}

	return nil
}

func (r *CategoryRepository) Delete(id, userID uuid.UUID) error {
	query := "DELETE FROM categories WHERE id = $1 AND user_id = $2"

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}
