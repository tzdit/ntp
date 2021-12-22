package models

import (
	"time"
)

//Instance model struct
type Instance struct {
	ID              string    `json:"id"`
	SeriesID        string    `json:"series_id"`
	SopUID          string    `json:"sop_uid"`
	SopCUID         string    `json:"sop_cuid"`
	InstanceNumber  int32     `json:"instance_number"`
	ContentDateTime time.Time `json:"content_datetime"`
	Archived        bool      `json:"archived"`
	Series          *Series   `json:"series"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedBy       string    `json:"deleted_by"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}
