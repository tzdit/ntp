package models

import (
	"time"
)

//LevelEntity DataStructure
type LevelEntity struct {
	ID          ID        `json:"id,omitempty"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	CreatedBy   int32     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
