package models

import (
	"time"
)

//ApplicationLog DataStructure
type ApplicationLog struct {
	ID        string    `json:"id,omitempty"`
	Level     string    `json:"level"`
	Host      string    `json:"host"`
	Process   string    `json:"process"`
	Message   string    `json:"message"`
	Exception string    `json:"exception"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at_at,omitempty"`
}
