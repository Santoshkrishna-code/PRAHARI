package recognition

import (
	"errors"
	"time"
)

// Recognition rewards positive safe behaviors.
type Recognition struct {
	ID                 string    `json:"id" db:"id"`
	ObservationID      string    `json:"observation_id" db:"observation_id"`
	RecognizedPersonID string    `json:"recognized_person_id" db:"recognized_person_id"`
	GrantedByID        string    `json:"granted_by_id" db:"granted_by_id"`
	Reason             string    `json:"reason" db:"reason"`
	GrantedAt          time.Time `json:"granted_at" db:"granted_at"`
}

// Validate checks domain invariants.
func (r *Recognition) Validate() error {
	if r.ObservationID == "" {
		return errors.New("observation ID reference is required")
	}
	if r.RecognizedPersonID == "" {
		return errors.New("recognized person ID is required")
	}
	return nil
}
