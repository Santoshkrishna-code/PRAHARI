package correctiveaction

import (
	"errors"
	"time"
)

// CorrectiveAction details CAPA corrective plans.
type CorrectiveAction struct {
	ID          string    `json:"id" db:"id"`
	FindingID   string    `json:"finding_id" db:"finding_id"`
	Description string    `json:"description" db:"description"`
	TargetDate  time.Time `json:"target_date" db:"target_date"`
	IsCompleted bool      `json:"is_completed" db:"is_completed"`
}

// Validate checks domain invariants.
func (ca *CorrectiveAction) Validate() error {
	if ca.FindingID == "" {
		return errors.New("finding ID reference is required")
	}
	if ca.Description == "" {
		return errors.New("action description cannot be empty")
	}
	return nil
}
