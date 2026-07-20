package schedule

import (
	"errors"
	"time"
)

// Frequency defines cycle rules.
type Frequency string

const (
	FreqDaily   Frequency = "DAILY"
	FreqWeekly  Frequency = "WEEKLY"
	FreqMonthly Frequency = "MONTHLY"
	FreqAnnual  Frequency = "ANNUAL"
)

// Schedule defines recurring safety audit rules.
type Schedule struct {
	ID                  string    `json:"id" db:"id"`
	TemplateID          string    `json:"template_id" db:"template_id"`
	Frequency           Frequency `json:"frequency" db:"frequency"`
	InspectorID         string    `json:"inspector_id" db:"inspector_id"`
	DepartmentID        string    `json:"department_id" db:"department_id"`
	LastExecutionDate   time.Time `json:"last_execution_date" db:"last_execution_date"`
	NextExecutionDate   time.Time `json:"next_execution_date" db:"next_execution_date"`
	IsActive            bool      `json:"is_active" db:"is_active"`
}

// Validate checks domain invariants.
func (s *Schedule) Validate() error {
	if s.TemplateID == "" {
		return errors.New("template ID is required for schedule")
	}
	if s.InspectorID == "" {
		return errors.New("inspector ID is required")
	}
	if s.NextExecutionDate.IsZero() {
		return errors.New("next execution date is required")
	}
	return nil
}

// CalculateNextDate advances scheduled triggers.
func (s *Schedule) CalculateNextDate() {
	s.LastExecutionDate = time.Now()
	switch s.Frequency {
	case FreqDaily:
		s.NextExecutionDate = s.LastExecutionDate.AddDate(0, 0, 1)
	case FreqWeekly:
		s.NextExecutionDate = s.LastExecutionDate.AddDate(0, 0, 7)
	case FreqMonthly:
		s.NextExecutionDate = s.LastExecutionDate.AddDate(0, 1, 0)
	case FreqAnnual:
		s.NextExecutionDate = s.LastExecutionDate.AddDate(1, 0, 0)
	}
}
