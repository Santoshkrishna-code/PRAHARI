package actionplan

import (
	"errors"
	"time"
)

// ActionPlan schedules corrective tasks.
type ActionPlan struct {
	ID           string    `json:"id" db:"id"`
	FindingID    string    `json:"finding_id" db:"finding_id"`
	Description  string    `json:"description" db:"description"`
	TargetDate   time.Time `json:"target_date" db:"target_date"`
	IsCompleted  bool      `json:"is_completed" db:"is_completed"`
}

// Validate checks domain invariants.
func (a *ActionPlan) Validate() error {
	if a.FindingID == "" {
		return errors.New("finding ID reference is required")
	}
	if a.Description == "" {
		return errors.New("action plan description cannot be empty")
	}
	return nil
}
