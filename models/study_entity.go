package models

import (
	"time"
)

//Study model struct
type Study struct {
	ID                      string    `json:"id" form:"id"`
	PatientID               string    `json:"patient_id" form:"patient_id"`
	PatientUUID             string    `json:"patient_uuid" form:"patient_uuid"`
	StudyInstanceUID        string    `json:"study_instance_uid" form:"study_instance_uid"`
	StudyDescription        string    `json:"study_description" form:"study_description"`
	StudyDate               string    `json:"study_date" form:"study_date"`
	StudyTime               string `json:"study_time,omitempty" form:"study_time"`
	NumberRelatedSeries     int32     `json:"number_related_series" form:"number_related_series"`
	NumberRelatedInstances  int32     `json:"number_related_instances" form:"number_related_instances"`
	ReferringPhysiciansName string    `json:"referring_physicians_name" form:"referring_physicians_name"`
	StudyNumber             string    `json:"study_number" form:"study_number"`
	AccessionNumber         string    `json:"accession_number" form:"accession_number"`
	ClinicalSummary         string    `json:"clinical_summary" form:"clinical_summary"`
	ReportedDate            string `json:"reported_datetime,omitempty" form:"reported_datetime"`
	ReportedStatus          bool      `json:"reported_status" form:"reported_status"`
	ReferralStatus          bool      `json:"referral_status" form:"referral_status"`
	Patient                 *Patient
	Series                  []*Series `json:"series"`
	CreatedBy               string    `json:"created_by"`
	CreatedAt               time.Time `json:"created_at,omitempty"`
	UpdatedBy               string    `json:"updated_by"`
	UpdatedAt               time.Time `json:"updated_at,omitempty"`
	DeletedBy               string    `json:"deleted_by"`
	DeletedAt               time.Time `json:"deleted_at,omitempty"`
}
