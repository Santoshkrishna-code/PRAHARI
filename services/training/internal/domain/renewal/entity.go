package renewal

import (
	"errors"
	"time"
)

// Renewal logs certification renewals schedules.
type Renewal struct {
	ID              string    `json:"id" db:"id"`
	CertificationID string    `json:"certification_id" db:"certification_id"`
	ScheduledDate   time.Time `json:"scheduled_date" db:"scheduled_date"`
	IsCompleted     bool      `json:"is_completed" db:"is_completed"`
}

// Validate checks domain invariants.
func (r *Renewal) Validate() error {
	if r.CertificationID == "" {
		return errors.New("certification ID reference is required")
	}
	return nil
}
