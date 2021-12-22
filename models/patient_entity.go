package models

import (
	"github.com/jackc/pgtype"
	"time"
)

//Patient DataStructure
type Patient struct {
	ID                 string      `json:"id,omitempty"`
	PatientMRN         string      `json:"pat_mrn"`
	PatientFirstName   string      `json:"pat_first_name"`
	PatientMiddleName  string      `json:"pat_middle_name"`
	PatientLastName    string      `json:"pat_last_name"`
	PatientBirthDate   pgtype.Date `json:"pat_birth_date"`
	PatientPhoneNumber string      `json:"pat_phone_number"`
	PatientAge         int32       `json:"pat_age" validate:"numeric"`
	PatientAddress     string      `json:"pat_address"`
	CreatedBy          string      `json:"created_by"`
	CreatedAt          time.Time   `json:"created_at,omitempty"`
	UpdatedBy          string      `json:"updated_by"`
	UpdatedAt          time.Time   `json:"updated_at,omitempty"`
	DeletedBy          string      `json:"deleted_by"`
	DeletedAt          time.Time   `json:"deleted_at,omitempty"`
}
