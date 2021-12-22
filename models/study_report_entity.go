package models

import (
	"time"
)

//StudyReport model struct
type StudyReport struct {
	ID                string    `json:"id,omitempty"`
	StudyID           string    `json:"study_id"`
	ReportedBy        string    `json:"reported_by"`
	ProcedureFindings string    `json:"procedure_findings"`
	Conclusion        string    `json:"conclusion"`
	ReportedDate      time.Time `json:"reported_date"`
	ReferralStatus    bool      `json:"referral_status"`
	CreatedBy         string    `json:"created_by"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedBy         string    `json:"updated_by"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
	DeletedBy         string    `json:"deleted_by"`
	DeletedAt         time.Time `json:"deleted_at,omitempty"`
}
