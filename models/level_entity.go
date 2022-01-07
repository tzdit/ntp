package models

import (
	"time"
)

//LevelEntity DataStructure
type LevelEntity struct {
	ID          ID        `json:"id,omitempty" form:"id"`
	Name        string    `json:"name" validate:"required" form:"name"`
	Description string    `json:"description" form:"description"`
	CreatedBy   int32     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
