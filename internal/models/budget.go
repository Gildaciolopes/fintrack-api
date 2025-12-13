package models

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID         uuid.UUID `json:"id" db:"id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id" binding:"required"`
	Amount     float64   `json:"amount" db:"amount" binding:"required,gt=0"`
	Month      time.Time `json:"month" db:"month" binding:"required"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Category   *Category `json:"category,omitempty" db:"-"`
}

type CreateBudgetRequest struct {
	CategoryID uuid.UUID `json:"category_id" binding:"required"`
	Amount     float64   `json:"amount" binding:"required,gt=0"`
	Month      time.Time `json:"month" binding:"required"`
}

type UpdateBudgetRequest struct {
	Amount float64   `json:"amount" binding:"omitempty,gt=0"`
	Month  time.Time `json:"month" binding:"omitempty"`
}

type BudgetWithSpent struct {
	Budget
	Spent      float64 `json:"spent"`
	Remaining  float64 `json:"remaining"`
	Percentage float64 `json:"percentage"`
}
