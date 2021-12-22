package models

import (
	"time"
)

//Ward DataStructure
type Ward struct {
	ID          string `json:"id,omitempty"`
	DistrictID  string `json:"district_id"`
	Name        string `json:"name"  validate:"required"`
	Code        string `json:"code"`
	Description string `json:"description"`
	District    *District
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedBy   string    `json:"deleted_by"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
