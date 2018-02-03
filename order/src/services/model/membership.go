package model

import (
	"time"
)

// Member model
type Member struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordSalt string    `json:"passwordSalt"`
	BirthDate    time.Time `json:"birthDate"`
	Version      int       `json:"version"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

//Members type
type Members []Member

// NewMember function for initialise Member model
func NewMember() *Member {
	now := time.Now()
	return &Member{
		Version:   0,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
