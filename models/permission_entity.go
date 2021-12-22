package models

import (
	"errors"
	"time"
)

var ErrInvalidPermissionEntity = errors.New("invalid permission")

// Permission Struct
type Permission struct {
	ID         string    `json:"id,omitempty"`
	Path       string    `json:"path,omitempty"`
	Method     string    `json:"method,omitempty"`
	Service    string    `json:"service,omitempty"`
	SubService string    `json:"sub_service,omitempty"`
	Action     string    `json:"action,omitempty"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedBy  string    `json:"deleted_by"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
}
