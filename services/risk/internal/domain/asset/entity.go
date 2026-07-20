package asset

import (
	"errors"
)

// Asset identifiers mapping.
type Asset struct {
	ID     string `json:"id" db:"id"`
	RiskID string `json:"risk_id" db:"risk_id"`
}

// Validate checks domain invariants.
func (a *Asset) Validate() error {
	if a.ID == "" {
		return errors.New("asset ID reference is required")
	}
	return nil
}
