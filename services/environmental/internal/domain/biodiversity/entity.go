package biodiversity

import (
	"errors"
	"time"
)

// BiodiversityRecord tracks conservation parameters.
type BiodiversityRecord struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	SpeciesName    string    `json:"species_name" db:"species_name"`
	Count          int       `json:"count" db:"count"`
	HealthStatus   string    `json:"health_status" db:"health_status"` // "HEALTHY", "THREATENED", "IMPACTED"
	ProtectedArea  bool      `json:"protected_area" db:"protected_area"`
	SurveyDate     time.Time `json:"survey_date" db:"survey_date"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks biodiversity attributes.
func (b *BiodiversityRecord) Validate() error {
	if b.PlantID == "" {
		return errors.New("plant ID is required")
	}
	if b.SpeciesName == "" {
		return errors.New("observed species name is required")
	}
	return nil
}
