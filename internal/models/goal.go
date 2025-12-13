package models

import (
	"time"

	"github.com/google/uuid"
)

type FinancialGoal struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	UserID        uuid.UUID  `json:"user_id" db:"user_id"`
	Title         string     `json:"title" db:"title" binding:"required,min=1,max=200"`
	TargetAmount  float64    `json:"target_amount" db:"target_amount" binding:"required,gt=0"`
	CurrentAmount float64    `json:"current_amount" db:"current_amount"`
	Deadline      *time.Time `json:"deadline" db:"deadline"`
	Status        string     `json:"status" db:"status" binding:"required,oneof=active completed cancelled"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateGoalRequest struct {
	Title         string     `json:"title" binding:"required,min=1,max=200"`
	TargetAmount  float64    `json:"target_amount" binding:"required,gt=0"`
	CurrentAmount float64    `json:"current_amount" binding:"omitempty,gte=0"`
	Deadline      *time.Time `json:"deadline"`
}

type UpdateGoalRequest struct {
	Title         string     `json:"title" binding:"omitempty,min=1,max=200"`
	TargetAmount  float64    `json:"target_amount" binding:"omitempty,gt=0"`
	CurrentAmount float64    `json:"current_amount" binding:"omitempty,gte=0"`
	Deadline      *time.Time `json:"deadline"`
	Status        string     `json:"status" binding:"omitempty,oneof=active completed cancelled"`
}

type ContributeGoalRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
