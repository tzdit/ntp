package models

import (
	"time"
)

//Ward DataStructure
type Ward struct {
	ID          string `json:"id,omitempty" form:"id"`
	DistrictId  string `json:"district_id" form:"district_id"`
	Name        string `json:"name"  validate:"required" form:"name"`
	Code        string `json:"code" form:"code"`
	Description string `json:"description" form:"description"`
	District    *District
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedBy   string    `json:"deleted_by"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
