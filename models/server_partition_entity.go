package models

import (
	"time"
)

//ServerPartitionEntity DataStructure
type ServerPartitionEntity struct {
	ID                int32
	StudyInstanceUid  string
	AeTitle           string
	Port              int16
	DefaultRemotePort int16
	StudyCount        int16
	InsertTime        time.Time
	LastAccessedTime  time.Time
	WriteLock         bool
	ReadLock          bool
	AcceptAnyDevice   bool
	AutoDeleteStudy   bool
	Enabled           bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}
