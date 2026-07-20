package schedule

import (
	"errors"
	"time"
)

// Schedule maps planning calendars.
type Schedule struct {
	ID                  string    `json:"id" db:"id"`
	MaintenanceID       string    `json:"maintenance_id" db:"maintenance_id"`
	ScheduledStartDate  time.Time `json:"scheduled_start_date" db:"scheduled_start_date"`
	ScheduledEndDate    time.Time `json:"scheduled_end_date" db:"scheduled_end_date"`
	EstimatedDowntimeMin int       `json:"estimated_downtime_min" db:"estimated_downtime_min"`
}

// Validate checks domain invariants.
func (s *Schedule) Validate() error {
	if s.MaintenanceID == "" {
		return errors.New("maintenance ID reference is required")
	}
	if s.ScheduledEndDate.Before(s.ScheduledStartDate) {
		return errors.New("end date must exceed scheduled start date")
	}
	return nil
}
