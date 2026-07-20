package reporter

import (
	"errors"
)

// Reporter profiles details.
type Reporter struct {
	ID          string `json:"id" db:"id"`
	NearMissID  string `json:"near_miss_id" db:"near_miss_id"`
	UserID      string `json:"user_id,omitempty" db:"user_id"`
	IsAnonymous bool   `json:"is_anonymous" db:"is_anonymous"`
}

// Validate checks domain invariants.
func (r *Reporter) Validate() error {
	if r.NearMissID == "" {
		return errors.New("near miss ID reference is required")
	}
	return nil
}
