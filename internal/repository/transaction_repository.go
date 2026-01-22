package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Gildaciolopes/fintrack-api/internal/models"
	"github.com/google/uuid"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	query := `
		INSERT INTO transactions (id, user_id, category_id, type, amount, description, date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	transaction.ID = uuid.New()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	return r.db.QueryRow(
		query,
		transaction.ID,
		transaction.UserID,
		transaction.CategoryID,
		transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.Date,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	).Scan(&transaction.ID, &transaction.CreatedAt, &transaction.UpdatedAt)
}

func (r *TransactionRepository) GetByID(id, userID uuid.UUID) (*models.Transaction, error) {
	query := `
		SELECT 
			t.id, t.user_id, t.category_id, t.type, t.amount, t.description, t.date, 
			t.created_at, t.updated_at,
			c.id, c.user_id, c.name, c.type, c.color, c.icon, c.created_at
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.id = $1 AND t.user_id = $2
	`

	transaction := &models.Transaction{}
	var category models.Category
	var categoryID, categoryUserID sql.NullString
	var categoryName, categoryType, categoryColor, categoryIcon sql.NullString
	var categoryCreatedAt sql.NullTime

	err := r.db.QueryRow(query, id, userID).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.CategoryID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Description,
		&transaction.Date,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&categoryID,
		&categoryUserID,
		&categoryName,
		&categoryType,
		&categoryColor,
		&categoryIcon,
		&categoryCreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("transaction not found")
	}
	if err != nil {
		return nil, err
	}

	if categoryID.Valid {
		categoryUUID, _ := uuid.Parse(categoryID.String)
		categoryUserUUID, _ := uuid.Parse(categoryUserID.String)
		category.ID = categoryUUID
		category.UserID = categoryUserUUID
		category.Name = categoryName.String
		category.Type = categoryType.String
		category.Color = categoryColor.String
		category.Icon = categoryIcon.String
		category.CreatedAt = categoryCreatedAt.Time
		transaction.Category = &category
	}

	return transaction, nil
}

func (r *TransactionRepository) GetAll(userID uuid.UUID, filters models.TransactionFilters) ([]models.Transaction, int64, error) {
	if filters.Page == 0 {
		filters.Page = 1
	}
	if filters.Limit == 0 {
		filters.Limit = 20
	}

	whereClause := "t.user_id = $1::uuid"
	args := []interface{}{userID}
	argPos := 2

	if filters.Type != "" {
		whereClause += fmt.Sprintf(" AND t.type = $%d", argPos)
		args = append(args, filters.Type)
		argPos++
	}

	if filters.CategoryID != nil {
		whereClause += fmt.Sprintf(" AND t.category_id = $%d::uuid", argPos)
		args = append(args, *filters.CategoryID)
		argPos++
	}

	if filters.StartDate != nil {
		whereClause += fmt.Sprintf(" AND t.date >= $%d::date", argPos)
		args = append(args, *filters.StartDate)
		argPos++
	}

	if filters.EndDate != nil {
		whereClause += fmt.Sprintf(" AND t.date <= $%d::date", argPos)
		args = append(args, *filters.EndDate)
		argPos++
	}

	if filters.MinAmount != nil {
		whereClause += fmt.Sprintf(" AND t.amount >= $%d", argPos)
		args = append(args, *filters.MinAmount)
		argPos++
	}

	if filters.MaxAmount != nil {
		whereClause += fmt.Sprintf(" AND t.amount <= $%d", argPos)
		args = append(args, *filters.MaxAmount)
		argPos++
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM transactions t WHERE %s", whereClause)
	var totalCount int64
	if err := r.db.QueryRow(countQuery, args...).Scan(&totalCount); err != nil {
		return nil, 0, err
	}

	offset := (filters.Page - 1) * filters.Limit
	query := fmt.Sprintf(`
		SELECT 
			t.id, t.user_id, t.category_id, t.type, t.amount, t.description, t.date, 
			t.created_at, t.updated_at,
			c.id, c.user_id, c.name, c.type, c.color, c.icon, c.created_at
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE %s
		ORDER BY t.date DESC, t.created_at DESC
		LIMIT $%d::int OFFSET $%d::int
	`, whereClause, argPos, argPos+1)

	args = append(args, filters.Limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		var category models.Category
		var categoryID, categoryUserID sql.NullString
		var categoryName, categoryType, categoryColor, categoryIcon sql.NullString
		var categoryCreatedAt sql.NullTime

		if err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.CategoryID,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Date,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&categoryID,
			&categoryUserID,
			&categoryName,
			&categoryType,
			&categoryColor,
			&categoryIcon,
			&categoryCreatedAt,
		); err != nil {
			return nil, 0, err
		}

		if categoryID.Valid {
			categoryUUID, _ := uuid.Parse(categoryID.String)
			categoryUserUUID, _ := uuid.Parse(categoryUserID.String)
			category.ID = categoryUUID
			category.UserID = categoryUserUUID
			category.Name = categoryName.String
			category.Type = categoryType.String
			category.Color = categoryColor.String
			category.Icon = categoryIcon.String
			category.CreatedAt = categoryCreatedAt.Time
			transaction.Category = &category
		}

		transactions = append(transactions, transaction)
	}

	return transactions, totalCount, rows.Err()
}

func (r *TransactionRepository) Update(id, userID uuid.UUID, updates map[string]interface{}) error {
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
		"UPDATE transactions SET %s WHERE id = $%d AND user_id = $%d",
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
		return fmt.Errorf("transaction not found")
	}

	return nil
}

func (r *TransactionRepository) Delete(id, userID uuid.UUID) error {
	query := "DELETE FROM transactions WHERE id = $1 AND user_id = $2"

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("transaction not found")
	}

	return nil
}

func (r *TransactionRepository) GetRecentTransactions(userID uuid.UUID, limit int) ([]models.Transaction, error) {
	query := `
		SELECT 
			t.id, t.user_id, t.category_id, t.type, t.amount, t.description, t.date, 
			t.created_at, t.updated_at,
			c.id, c.user_id, c.name, c.type, c.color, c.icon, c.created_at
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.user_id = $1
		ORDER BY t.date DESC, t.created_at DESC
		LIMIT $2
	`

	rows, err := r.db.Query(query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		var category models.Category
		var categoryID, categoryUserID sql.NullString
		var categoryName, categoryType, categoryColor, categoryIcon sql.NullString
		var categoryCreatedAt sql.NullTime

		if err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.CategoryID,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Date,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&categoryID,
			&categoryUserID,
			&categoryName,
			&categoryType,
			&categoryColor,
			&categoryIcon,
			&categoryCreatedAt,
		); err != nil {
			return nil, err
		}

		if categoryID.Valid {
			categoryUUID, _ := uuid.Parse(categoryID.String)
			categoryUserUUID, _ := uuid.Parse(categoryUserID.String)
			category.ID = categoryUUID
			category.UserID = categoryUserUUID
			category.Name = categoryName.String
			category.Type = categoryType.String
			category.Color = categoryColor.String
			category.Icon = categoryIcon.String
			category.CreatedAt = categoryCreatedAt.Time
			transaction.Category = &category
		}

		transactions = append(transactions, transaction)
	}

	return transactions, rows.Err()
}
