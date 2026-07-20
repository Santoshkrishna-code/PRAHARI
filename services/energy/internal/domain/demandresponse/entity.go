package demandresponse

import (
	"errors"
	"time"
)

// Event tracks peak load shedding programs.
type Event struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	TargetShedKW   float64   `json:"target_shed_kw" db:"target_shed_kw"`
	ActualShedKW   float64   `json:"actual_shed_kw" db:"actual_shed_kw"`
	StartTime      time.Time `json:"start_time" db:"start_time"`
	EndTime        time.Time `json:"end_time" db:"end_time"`
	IsSuccessful   bool      `json:"is_successful" db:"is_successful"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks events.
func (e *Event) Validate() error {
	if e.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}
