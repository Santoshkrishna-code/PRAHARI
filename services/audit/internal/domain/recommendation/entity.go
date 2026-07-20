package recommendation

import (
	"errors"
)

// Recommendation details improvements.
type Recommendation struct {
	ID          string `json:"id" db:"id"`
	FindingID   string `json:"finding_id" db:"finding_id"`
	Suggestion  string `json:"suggestion" db:"suggestion"`
}

// Validate checks domain invariants.
func (r *Recommendation) Validate() error {
	if r.FindingID == "" {
		return errors.New("finding ID reference is required")
	}
	if r.Suggestion == "" {
		return errors.New("suggestion text cannot be empty")
	}
	return nil
}
