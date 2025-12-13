package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id"`
	CategoryID  *uuid.UUID `json:"category_id" db:"category_id"`
	Type        string     `json:"type" db:"type" binding:"required,oneof=income expense"`
	Amount      float64    `json:"amount" db:"amount" binding:"required,gt=0"`
	Description *string    `json:"description" db:"description"`
	Date        time.Time  `json:"date" db:"date" binding:"required"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	Category    *Category  `json:"category,omitempty" db:"-"`
}

type CreateTransactionRequest struct {
	CategoryID  *uuid.UUID `json:"category_id"`
	Type        string     `json:"type" binding:"required,oneof=income expense"`
	Amount      float64    `json:"amount" binding:"required,gt=0"`
	Description *string    `json:"description"`
	Date        time.Time  `json:"date" binding:"required"`
}

type UpdateTransactionRequest struct {
	CategoryID  *uuid.UUID `json:"category_id"`
	Type        string     `json:"type" binding:"omitempty,oneof=income expense"`
	Amount      float64    `json:"amount" binding:"omitempty,gt=0"`
	Description *string    `json:"description"`
	Date        time.Time  `json:"date"`
}

type TransactionFilters struct {
	Type       string     `form:"type" binding:"omitempty,oneof=income expense"`
	CategoryID *uuid.UUID `form:"category_id"`
	StartDate  *time.Time `form:"start_date"`
	EndDate    *time.Time `form:"end_date"`
	MinAmount  *float64   `form:"min_amount" binding:"omitempty,gte=0"`
	MaxAmount  *float64   `form:"max_amount" binding:"omitempty,gte=0"`
	Page       int        `form:"page" binding:"omitempty,gte=1"`
	Limit      int        `form:"limit" binding:"omitempty,gte=1,lte=100"`
}
