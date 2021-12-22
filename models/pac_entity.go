package models

import (
	"time"
)

//Pac model struct
type Pac struct {
	ID              string    `json:"id,omitempty"`
	HospitalID      string    `json:"hospital_id"`
	Name            string    `json:"name"`
	IpAddress       string    `json:"ip_address"`
	Port            int32     `json:"port"`
	AllowRecv       bool      `json:"allow_recv"`
	AllowSend       bool      `json:"allow_send"`
	Description     string    `json:"description"`
	LastAccessTime  time.Time `json:"last_access_time"`
	AcceptAnyDevice bool      `json:"accept_any_device"`
	AutoDeleteStudy bool      `json:"auto_delete_study"`
	Enabled         bool      `json:"enabled"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedBy       string    `json:"deleted_by"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}
