package models

import (
	"time"
)

//Region DataStructure
type Region struct {
	ID        string    `json:"id,omitempty" form:"id"`
	Name      string    `json:"name"  validate:"required" form:"name"`
	Code      string    `json:"code" validate:"required" form:"code"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedBy string    `json:"deleted_by"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
