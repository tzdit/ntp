package models

import (
	"errors"
	"time"
)

var ErrInvalidUserCategoryEntity = errors.New("invalid user category")

// UserCategory DatasStructure
type UserCategory struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy string
}
