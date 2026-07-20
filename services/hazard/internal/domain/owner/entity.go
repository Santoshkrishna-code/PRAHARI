package owner

import (
	"errors"
)

// Owner assigns specific personnel roles.
type Owner struct {
	ID           string `json:"id" db:"id"`
	HazardID     string `json:"hazard_id" db:"hazard_id"`
	UserID       string `json:"user_id" db:"user_id"`
	IsLeadAssign bool   `json:"is_lead_assign" db:"is_lead_assign"`
}

// Validate checks domain invariants.
func (o *Owner) Validate() error {
	if o.HazardID == "" {
		return errors.New("hazard ID reference is required")
	}
	if o.UserID == "" {
		return errors.New("user ID reference is required")
	}
	return nil
}
