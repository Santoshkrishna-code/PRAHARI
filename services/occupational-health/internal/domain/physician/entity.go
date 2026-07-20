package physician

import (
	"errors"
	"time"
)

// Physician represents clinic medical doctors verifying health states.
type Physician struct {
	ID             string    `json:"id" db:"id"`
	LicenseNumber  string    `json:"license_number" db:"license_number"`
	FullName       string    `json:"full_name" db:"full_name"`
	Specialty      string    `json:"specialty" db:"specialty"` // e.g. "OCCUPATIONAL_MEDICINE"
	ContactEmail   string    `json:"contact_email" db:"contact_email"`
	ClinicID       string    `json:"clinic_id" db:"clinic_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks physician attributes.
func (p *Physician) Validate() error {
	if p.LicenseNumber == "" {
		return errors.New("license number is required")
	}
	if p.FullName == "" {
		return errors.New("full name is required")
	}
	return nil
}
