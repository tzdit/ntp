package models

import (
	"time"
)

//HospitalPatient model struct
type HospitalPatient struct {
	ID         string    `json:"id,omitempty" form:"id"`
	HospitalID string    `json:"hospital_id" form:"hospital_id"`
	PatientID  string    `json:"patient_id" form:"patient_id"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedBy  string    `json:"deleted_by"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
}
