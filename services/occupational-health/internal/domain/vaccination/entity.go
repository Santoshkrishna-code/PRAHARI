package vaccination

import (
	"errors"
	"time"
)

// Vaccination tracks workforce immunizations and booster timelines.
type Vaccination struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	VaccineName     string    `json:"vaccine_name" db:"vaccine_name"` // e.g. "Tetanus", "Hepatitis B", "COVID-19"
	DoseNumber      int       `json:"dose_number" db:"dose_number"`
	AdministeredDate time.Time `json:"administered_date" db:"administered_date"`
	ExpiryDate      time.Time `json:"expiry_date" db:"expiry_date"`
	BatchNumber     string    `json:"batch_number" db:"batch_number"`
	ProviderName    string    `json:"provider_name" db:"provider_name"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks vaccination dates.
func (v *Vaccination) Validate() error {
	if v.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if v.VaccineName == "" {
		return errors.New("vaccine name is required")
	}
	return nil
}
