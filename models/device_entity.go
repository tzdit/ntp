package models

import (
	"time"
)

//Device DataStructure
type Device struct {
	ID             string    `json:"id,omitempty"`
	DeviceTypeId   string    `json:"device_type_id"`
	DeviceTypeName string    `json:"device_type_name"`
	HospitalId     string    `json:"hospital_id"`
	HospitalName   string    `json:"hospital_name"`
	Name           string    `json:"name" validate:"required"`
	AeTitle        string    `json:"ae_title" validate:"required"`
	IpAddress      string    `json:"ip_address" validate:"required"`
	Port           int16     `json:"port"`
	AllowStorage   bool      `json:"allow_storage"`
	AllowRetrieve  bool      `json:"allow_retrieve"`
	Description    string    `json:"description"`
	LastAccessTime time.Time `json:"last_access_time"`
	DeviceType     *DeviceType
	Hospital       *Hospital
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	DeletedBy      string    `json:"deleted_by"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
}
