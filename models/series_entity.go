package models

import (
	"time"
)

//Series model struct
type Series struct {
	ID                           string    `json:"id" form:"id"`
	StudyID                      string    `json:"study_id" form:"study_id"`
	SeriesNumber                 string    `json:"series_number" form:"series_number"`
	Modality                     string    `json:"modality" form:"modality"`
	SeriesInstanceUID            string    `json:"series_instance_uid" form:"series_instance_uid"`
	SeriesDescription            string    `json:"series_description" form:"series_description"`
	NumberRelatedSeriesInstances int32     `json:"number_related_series_instances" form:"number_related_series_instances"`
	SourceApplicationEntityTitle string    `json:"source_applicatiion_entity_title" form:"source_applicatiion_entity_title"`
	Availability                 bool      `json:"availability" form:"availability"`
	Study                        *Study    `json:"study"`
	CreatedBy                    string    `json:"created_by"`
	CreatedAt                    time.Time `json:"created_at,omitempty"`
	UpdatedBy                    string    `json:"updated_by"`
	UpdatedAt                    time.Time `json:"updated_at,omitempty"`
	DeletedBy                    string    `json:"deleted_by"`
	DeletedAt                    time.Time `json:"deleted_at,omitempty"`
}
