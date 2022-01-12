package models

import (
	"time"
)

//Hospital DataStructure
type Hospital struct {
	ID              string  `json:"id,omitempty" form:"id"`
	WardID          string  `json:"ward_id" form:"ward_id"`
	HospitalLevelID string  `json:"hospital_level_id" form:"hospital_level_id"`
	HospitalTypeID  string  `json:"hospital_type_id" form:"hospital_type_id"`
	HospitalOwnerID string  `json:"hospital_owner_id" form:"hospital_owner_id"`
	HospitalNumber  string  `json:"hospital_number" form:"hospital_number"`
	HospitalName    string  `json:"hospital_name" form:"hospital_name"`
	Longitude       float64 `json:"longitude" form:"longitude"`
	Latitude        float64 `json:"latitude" form:"latitude"`
	Description     string  `json:"description" form:"description"`
	Ward            *Ward
	HospitalLevel   *HospitalLevel
	HospitalType    *HospitalType
	HospitalOwner   *HospitalOwner
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedBy       string    `json:"deleted_by"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}
