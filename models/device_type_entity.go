package models

import (
	"time"
)

//DeviceType DataStructure
type DeviceType struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedBy   string    `json:"deleted_by"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
