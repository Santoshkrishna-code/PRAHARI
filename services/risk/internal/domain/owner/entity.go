package owner

import (
	"errors"
)

// Owner assigns risk ownership.
type Owner struct {
	ID       string `json:"id" db:"id"`
	RiskID   string `json:"risk_id" db:"risk_id"`
	UserID   string `json:"user_id" db:"user_id"`
	RoleType string `json:"role_type" db:"role_type"` // e.g. Lead Assessor, Safety Supervisor
}

// Validate checks domain invariants.
func (o *Owner) Validate() error {
	if o.RiskID == "" {
		return errors.New("risk ID reference is required")
	}
	if o.UserID == "" {
		return errors.New("user ID reference is required")
	}
	return nil
}
