package models

import "time"

//Sync model struct
type Sync struct {
	ID           string    `json:"id,omitempty"`
	TableName    string    `json:"table_name"`
	LastSyncTime time.Time `json:"last_sync_time"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}
