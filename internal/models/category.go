package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name" binding:"required,min=1,max=100"`
	Type      string    `json:"type" db:"type" binding:"required,oneof=income expense"`
	Color     string    `json:"color" db:"color" binding:"required,hexcolor"`
	Icon      string    `json:"icon" db:"icon" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateCategoryRequest struct {
	Name  string `json:"name" binding:"required,min=1,max=100"`
	Type  string `json:"type" binding:"required,oneof=income expense"`
	Color string `json:"color" binding:"required,hexcolor"`
	Icon  string `json:"icon" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name  string `json:"name" binding:"omitempty,min=1,max=100"`
	Type  string `json:"type" binding:"omitempty,oneof=income expense"`
	Color string `json:"color" binding:"omitempty,hexcolor"`
	Icon  string `json:"icon" binding:"omitempty"`
}
