package models

import (
	"time"
)

//Device DataStructure
type Device struct {
	ID             string    `json:"id,omitempty" form:"id"`
	DeviceTypeId   string    `json:"device_type_id" form:"device_type_id"`
	DeviceTypeName string    `json:"device_type_name" form:"device_type_name"`
	HospitalId     string    `json:"hospital_id" form:"hospital_id"`
	HospitalName   string    `json:"hospital_name" form:"hospital_name"`
	Name           string    `json:"name" validate:"required" form:"name"`
	AeTitle        string    `json:"ae_title" validate:"required" form:"ae_title"`
	IpAddress      string    `json:"ip_address" validate:"required" form:"ip_address"`
	Port           int16     `json:"port" form:"port"`
	AllowStorage   bool      `json:"allow_storage" form:"allow_storage"`
	AllowRetrieve  bool      `json:"allow_retrieve" form:"allow_retrieve"`
	Description    string    `json:"description" form:"description"`
	LastAccessTime time.Time `json:"last_access_time" form:"last_access_time"`
	DeviceType     *DeviceType
	Hospital       *Hospital
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	DeletedBy      string    `json:"deleted_by"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
}
