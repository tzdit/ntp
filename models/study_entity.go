package models

import (
	"github.com/jackc/pgtype"
	"time"
)

//Study model struct
type Study struct {
	ID                      string      `json:"id"`
	PatientID               string      `json:"patient_id"`
	PatientUUID             string      `json:"patient_uuid"`
	StudyInstanceUID        string      `json:"study_instance_uid"`
	StudyDescription        string      `json:"study_description"`
	StudyDate               pgtype.Date `json:"study_date"`
	StudyTime               time.Time   `json:"study_time,omitempty"`
	NumberRelatedSeries     int32       `json:"number_related_series"`
	NumberRelatedInstances  int32       `json:"number_related_instances"`
	ReferringPhysiciansName string      `json:"referring_physicians_name"`
	StudyNumber             string      `json:"study_number"`
	AccessionNumber         string      `json:"accession_number"`
	ClinicalSummary         string      `json:"clinical_summary"`
	ReportedDate            time.Time   `json:"reported_datetime,omitempty"`
	ReportedStatus          bool        `json:"reported_status"`
	ReferralStatus          bool        `json:"referral_status"`
	Patient                 *Patient
	Series                  []*Series `json:"series"`
	CreatedBy               string    `json:"created_by"`
	CreatedAt               time.Time `json:"created_at,omitempty"`
	UpdatedBy               string    `json:"updated_by"`
	UpdatedAt               time.Time `json:"updated_at,omitempty"`
	DeletedBy               string    `json:"deleted_by"`
	DeletedAt               time.Time `json:"deleted_at,omitempty"`
}
