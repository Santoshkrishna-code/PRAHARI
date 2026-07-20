package observation

import (
	"errors"
	"time"
)

// Observation tracks safety notes recorded on inspections.
type Observation struct {
	ID           string    `json:"id" db:"id"`
	InspectionID string    `json:"inspection_id" db:"inspection_id"`
	Description  string    `json:"description" db:"description"`
	ObserverID   string    `json:"observer_id" db:"observer_id"`
	ObservedAt   time.Time `json:"observed_at" db:"observed_at"`
}

// Validate checks domain invariants for Observation.
func (o *Observation) Validate() error {
	if o.InspectionID == "" {
		return errors.New("inspection ID is required for observation")
	}
	if o.Description == "" {
		return errors.New("observation description is required")
	}
	if o.ObserverID == "" {
		return errors.New("observer user ID is required")
	}
	return nil
}
