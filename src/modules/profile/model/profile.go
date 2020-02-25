package model

import (
	"time"
)

// Profile struct
type Profile struct {
	ID string
	FirsName string
	LastName string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time

}

// Profiles type profile list
type Profiles [] Profile

// NewProfile Profile's Construct

func NewProfile () *Profile{
	return &Profile{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

