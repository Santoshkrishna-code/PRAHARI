package observedperson

import (
	"errors"
)

// ObservedPerson maps target profile details.
type ObservedPerson struct {
	ID            string `json:"id" db:"id"`
	ObservationID string `json:"observation_id" db:"observation_id"`
	UserID        string `json:"user_id,omitempty" db:"user_id"`
	IsAnonymous   bool   `json:"is_anonymous" db:"is_anonymous"`
}

// Validate checks domain invariants.
func (op *ObservedPerson) Validate() error {
	if op.ObservationID == "" {
		return errors.New("observation ID reference is required")
	}
	return nil
}
