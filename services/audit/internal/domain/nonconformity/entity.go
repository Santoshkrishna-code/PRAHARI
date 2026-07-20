package nonconformity

import (
	"errors"
)

// NonConformity logs major EHS safety violations.
type NonConformity struct {
	ID          string `json:"id" db:"id"`
	FindingID   string `json:"finding_id" db:"finding_id"`
	Severity    string `json:"severity" db:"severity"` // Minor, Major
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (nc *NonConformity) Validate() error {
	if nc.FindingID == "" {
		return errors.New("finding ID reference is required")
	}
	if nc.Description == "" {
		return errors.New("NCR description cannot be empty")
	}
	return nil
}
