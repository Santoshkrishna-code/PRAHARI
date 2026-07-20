package coaching

import (
	"errors"
	"time"
)

// CoachingSession records safety coaching feedback dialogues.
type CoachingSession struct {
	ID           string    `json:"id" db:"id"`
	ObservationID string    `json:"observation_id" db:"observation_id"`
	CoachID      string    `json:"coach_id" db:"coach_id"`
	SessionDate  time.Time `json:"session_date" db:"session_date"`
	Topics       string    `json:"topics" db:"topics"`
	Feedback     string    `json:"feedback" db:"feedback"`
}

// Validate checks domain invariants.
func (cs *CoachingSession) Validate() error {
	if cs.ObservationID == "" {
		return errors.New("observation ID reference is required")
	}
	if cs.CoachID == "" {
		return errors.New("coach ID reference is required")
	}
	return nil
}
