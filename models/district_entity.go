package models

import (
	"time"
)

//District DataStructure
type District struct {
	ID          string `json:"id,omitempty"`
	RegionID    string `json:"region_id"`
	Name        string `json:"name"  validate:"required"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Region      *Region
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedBy   string    `json:"deleted_by"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
