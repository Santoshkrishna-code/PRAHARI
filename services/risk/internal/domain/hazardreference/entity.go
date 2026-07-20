package hazardreference

import (
	"errors"
)

// HazardReference links hazard safety logs.
type HazardReference struct {
	ID         string `json:"id" db:"id"`
	RiskID     string `json:"risk_id" db:"risk_id"`
	HazardID   string `json:"hazard_id" db:"hazard_id"`
	LinkReason string `json:"link_reason,omitempty" db:"link_reason"`
}

// Validate checks domain invariants.
func (hr *HazardReference) Validate() error {
	if hr.RiskID == "" {
		return errors.New("risk ID reference is required")
	}
	if hr.HazardID == "" {
		return errors.New("hazard ID reference is required")
	}
	return nil
}
