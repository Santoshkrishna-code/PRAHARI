package observer

import (
	"errors"
)

// Observer maps person logging safety observations.
type Observer struct {
	ID            string `json:"id" db:"id"`
	ObservationID string `json:"observation_id" db:"observation_id"`
	UserID        string `json:"user_id" db:"user_id"`
	IsAnonymous   bool   `json:"is_anonymous" db:"is_anonymous"`
}

// Validate checks domain invariants.
func (o *Observer) Validate() error {
	if o.ObservationID == "" {
		return errors.New("observation ID reference is required")
	}
	return nil
}
