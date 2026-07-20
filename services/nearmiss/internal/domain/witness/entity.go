package witness

import (
	"errors"
)

// Witness maps crew member declarations.
type Witness struct {
	ID         string `json:"id" db:"id"`
	NearMissID string `json:"near_miss_id" db:"near_miss_id"`
	UserID     string `json:"user_id" db:"user_id"`
	Statement  string `json:"statement" db:"statement"`
}

// Validate checks domain invariants.
func (w *Witness) Validate() error {
	if w.NearMissID == "" {
		return errors.New("near miss ID reference is required")
	}
	if w.UserID == "" {
		return errors.New("user ID reference is required")
	}
	return nil
}
