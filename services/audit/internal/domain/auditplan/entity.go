package auditplan

import (
	"errors"
	"time"
)

// AuditPlan defines audit timeline schedules.
type AuditPlan struct {
	ID             string    `json:"id" db:"id"`
	AuditProgramID string    `json:"audit_program_id" db:"audit_program_id"`
	ScheduledStart time.Time `json:"scheduled_start" db:"scheduled_start"`
	ScheduledEnd   time.Time `json:"scheduled_end" db:"scheduled_end"`
}

// Validate checks domain invariants.
func (ap *AuditPlan) Validate() error {
	if ap.AuditProgramID == "" {
		return errors.New("audit program ID reference is required")
	}
	if ap.ScheduledEnd.Before(ap.ScheduledStart) {
		return errors.New("scheduled end date must be after start date")
	}
	return nil
}
