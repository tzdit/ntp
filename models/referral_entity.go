package models

import (
	"time"
)

//Referral model struct
type Referral struct {
	ID                 string    `json:"id,omitempty" form:"id"`
	StudyID            string    `json:"study_id" form:"study_id"`
	SenderHospitalID   string    `json:"sender_hospital_id" form:"sender_hospital_id"`
	ReceiverHospitalID string    `json:"receiver_hospital_id" form:"receiver_hospital_id"`
	Priority           string    `json:"priority" form:"priority"`
	Description        string    `json:"description" form:"description"`
	CreatedBy          string    `json:"created_by"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedBy          string    `json:"updated_by"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	DeletedBy          string    `json:"deleted_by"`
	DeletedAt          time.Time `json:"deleted_at,omitempty"`
}
