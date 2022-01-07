package models

import (
	"errors"
	"time"
)

var ErrInvalidPermissionEntity = errors.New("invalid permission")

// Permission Struct
type Permission struct {
	ID         string    `json:"id,omitempty" form:"id"`
	Path       string    `json:"path,omitempty" form:"path"`
	Method     string    `json:"method,omitempty" form:"method"`
	Service    string    `json:"service,omitempty" form:"service"`
	SubService string    `json:"sub_service,omitempty" form:"sub_service"`
	Action     string    `json:"action,omitempty" form:"action"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedBy  string    `json:"deleted_by"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
}
