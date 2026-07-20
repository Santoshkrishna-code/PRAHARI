package observation

import (
	"errors"
)

// Observation details.
type Observation struct {
	ID          string `json:"id" db:"id"`
	FindingID   string `json:"finding_id" db:"finding_id"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (o *Observation) Validate() error {
	if o.FindingID == "" {
		return errors.New("finding ID reference is required")
	}
	if o.Description == "" {
		return errors.New("observation description cannot be empty")
	}
	return nil
}
