package followup

import (
	"errors"
	"time"
)

// FollowUp checks coaching results feedback loop.
type FollowUp struct {
	ID           string    `json:"id" db:"id"`
	ObservationID string    `json:"observation_id" db:"observation_id"`
	FollowerID   string    `json:"follower_id" db:"follower_id"`
	FollowUpDate time.Time `json:"follow_up_date" db:"follow_up_date"`
	Notes        string    `json:"notes" db:"notes"`
	IsPassed     bool      `json:"is_passed" db:"is_passed"`
}

// Validate checks domain invariants.
func (f *FollowUp) Validate() error {
	if f.ObservationID == "" {
		return errors.New("observation ID is required")
	}
	if f.FollowerID == "" {
		return errors.New("follower ID is required")
	}
	return nil
}
