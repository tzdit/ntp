package models

import (
	"time"
)

//Hospital DataStructure
type Hospital struct {
	ID              string  `json:"id,omitempty"`
	WardId          string  `json:"ward_id"`
	HospitalLevelId string  `json:"hospital_level_id"`
	HospitalTypeId  string  `json:"hospital_type_id"`
	HospitalOwnerId string  `json:"hospital_owner_id"`
	HospitalNumber  string  `json:"hospital_number"`
	HospitalName    string  `json:"hospital_name"`
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
	Description     string  `json:"description"`
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
