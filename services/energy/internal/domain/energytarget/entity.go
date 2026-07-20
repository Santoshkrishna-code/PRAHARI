package energytarget

import (
	"errors"
	"time"
)

// Target maps conservation targets.
type Target struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	Title          string    `json:"title" db:"title"` // e.g. "Reduce grid electricity use by 5%"
	TargetValue    float64   `json:"target_value" db:"target_value"`
	CurrentValue   float64   `json:"current_value" db:"current_value"`
	UnitOfMeasure  string    `json:"unit_of_measure" db:"unit_of_measure"`
	Deadline       time.Time `json:"deadline" db:"deadline"`
	Status         string    `json:"status" db:"status"` // "ACTIVE", "EXCEEDED", "ACHIEVED"
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks targets.
func (t *Target) Validate() error {
	if t.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if t.Title == "" {
		return errors.New("target title description is required")
	}
	return nil
}
