package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Email     string     `json:"email" db:"email"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type AuthUser struct {
	ID    uuid.UUID `json:"sub"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
	Exp   int64     `json:"exp"`
}
