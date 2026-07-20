package environmentalobjective

import (
	"errors"
	"time"
)

// Objective structures sustainability parameters.
type Objective struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	Title          string    `json:"title" db:"title"` // e.g. "Reduce Carbon footprint by 10%"
	TargetValue    float64   `json:"target_value" db:"target_value"`
	CurrentValue   float64   `json:"current_value" db:"current_value"`
	UnitOfMeasure  string    `json:"unit_of_measure" db:"unit_of_measure"`
	Deadline       time.Time `json:"deadline" db:"deadline"`
	Status         string    `json:"status" db:"status"` // "ON_TRACK", "DELAYED", "ACHIEVED"
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks target metrics.
func (o *Objective) Validate() error {
	if o.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if o.Title == "" {
		return errors.New("objective title is required")
	}
	return nil
}
