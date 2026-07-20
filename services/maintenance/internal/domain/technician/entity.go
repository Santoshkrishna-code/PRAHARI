package technician

import (
	"errors"
)

// Technician maps credentials profiles.
type Technician struct {
	ID            string `json:"id" db:"id"`
	MaintenanceID string `json:"maintenance_id" db:"maintenance_id"`
	UserID        string `json:"user_id" db:"user_id"`
	LeadTechnician bool  `json:"lead_technician" db:"lead_technician"`
}

// Validate checks domain invariants.
func (t *Technician) Validate() error {
	if t.MaintenanceID == "" {
		return errors.New("maintenance ID reference is required")
	}
	if t.UserID == "" {
		return errors.New("user ID reference is required")
	}
	return nil
}
