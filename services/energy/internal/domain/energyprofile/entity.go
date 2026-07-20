package energyprofile

import (
	"errors"
	"time"
)

// Profile registers main business facilities energy profiles.
type Profile struct {
	ID           string    `json:"id" db:"id"`
	PlantID      string    `json:"plant_id" db:"plant_id"`
	DepartmentID string    `json:"department_id" db:"department_id"`
	FacilityName string    `json:"facility_name" db:"facility_name"`
	TargetScore  float64   `json:"target_score" db:"target_score"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks profile values.
func (p *Profile) Validate() error {
	if p.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if p.FacilityName == "" {
		return errors.New("facility name is required")
	}
	return nil
}
