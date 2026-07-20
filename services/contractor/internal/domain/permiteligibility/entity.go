package permiteligibility

import (
	"errors"
)

// PermitEligibility validates contractor permit qualification criteria.
type PermitEligibility struct {
	ID              string `json:"id" db:"id"`
	WorkerID        string `json:"worker_id" db:"worker_id"`
	IsEligible      bool   `json:"is_eligible" db:"is_eligible"`
	IneligibilityReason string `json:"ineligibility_reason,omitempty" db:"ineligibility_reason"`
}

// Validate checks domain invariants.
func (pe *PermitEligibility) Validate() error {
	if pe.WorkerID == "" {
		return errors.New("worker ID is required")
	}
	return nil
}
