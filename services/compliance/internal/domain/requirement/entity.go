package requirement

import (
	"errors"
)

// Requirement details rules clauses.
type Requirement struct {
	ID          string `json:"id" db:"id"`
	ObligationID string `json:"obligation_id" db:"obligation_id"`
	Clause      string `json:"clause" db:"clause"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (r *Requirement) Validate() error {
	if r.ObligationID == "" {
		return errors.New("obligation ID reference is required")
	}
	if r.Clause == "" {
		return errors.New("requirement clause is required")
	}
	return nil
}
