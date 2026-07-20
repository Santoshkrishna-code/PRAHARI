package returntowork

import (
	"errors"
	"time"
)

// Program tracks returning plans for workers after medical incidents.
type Program struct {
	ID                 string    `json:"id" db:"id"`
	HealthProfileID    string    `json:"health_profile_id" db:"health_profile_id"`
	TargetReturnDate   time.Time `json:"target_return_date" db:"target_return_date"`
	ActualReturnDate   time.Time `json:"actual_return_date" db:"actual_return_date"`
	IsPhased           bool      `json:"is_phased" db:"is_phased"`
	WeeklyTargetHours  int       `json:"weekly_target_hours" db:"weekly_target_hours"`
	Status             string    `json:"status" db:"status"` // "IN_PROGRESS", "SUCCESSFUL", "FAILED"
	Notes              string    `json:"notes" db:"notes"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks entity values.
func (p *Program) Validate() error {
	if p.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	return nil
}
