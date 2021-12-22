package models

import (
	"time"
)

//User defines class user data
type User struct {
	ID              string    `json:"id,omitempty"`
	HospitalID      string    `json:"hospital_id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email,omitempty"`
	EmailVerifiedAt time.Time `json:"email_verified_at,omitempty"`
	Password        string    `json:"password,omitempty"`
	LoginAttempt    int32     `json:"login_attempt,omitempty"`
	RememberToken   string    `json:"remember_token,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}
