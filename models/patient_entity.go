package models

import (
	"time"
)

//Patient DataStructure
type Patient struct {
	ID                 string      `json:"id,omitempty" form:"id"`
	PatientMRN         string      `json:"pat_mrn" form:"pat_mrn"`
	PatientFirstName   string      `json:"pat_first_name" form:"pat_first_name"`
	PatientMiddleName  string      `json:"pat_middle_name" form:"pat_middle_name"`
	PatientLastName    string      `json:"pat_last_name" form:"pat_last_name"`
	PatientBirthDate   string `json:"pat_birth_date" form:"pat_birth_date"`
	PatientPhoneNumber string      `json:"pat_phone_number" form:"pat_phone_number"`
	PatientAge         int32       `json:"pat_age" validate:"numeric" form:"pat_age"`
	PatientAddress     string      `json:"pat_address" form:"pat_address"`
	CreatedBy          string      `json:"created_by"`
	CreatedAt          time.Time   `json:"created_at,omitempty"`
	UpdatedBy          string      `json:"updated_by"`
	UpdatedAt          time.Time   `json:"updated_at,omitempty"`
	DeletedBy          string      `json:"deleted_by"`
	DeletedAt          time.Time   `json:"deleted_at,omitempty"`
}
