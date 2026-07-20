package laboratory

import (
	"errors"
	"time"
)

// Laboratory registers environmental labs.
type Laboratory struct {
	ID             string    `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	LicenseNumber  string    `json:"license_number" db:"license_number"`
	ContactEmail   string    `json:"contact_email" db:"contact_email"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks lab.
func (l *Laboratory) Validate() error {
	if l.Name == "" {
		return errors.New("laboratory name is required")
	}
	return nil
}
