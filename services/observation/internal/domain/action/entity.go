package action

import (
	"errors"
	"time"
)

// Action details corrective tasks.
type Action struct {
	ID                 string    `json:"id" db:"id"`
	ObservationID      string    `json:"observation_id" db:"observation_id"`
	Description        string    `json:"description" db:"description"`
	TargetDate         time.Time `json:"target_date" db:"target_date"`
	ResponsiblePartyID string    `json:"responsible_party_id" db:"responsible_party_id"`
	IsImplemented      bool      `json:"is_implemented" db:"is_implemented"`
}

// Validate checks domain invariants.
func (a *Action) Validate() error {
	if a.ObservationID == "" {
		return errors.New("observation ID reference is required")
	}
	if a.Description == "" {
		return errors.New("action description cannot be empty")
	}
	return nil
}
