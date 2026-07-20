package verification

import (
	"errors"
	"time"
)

// Verification records post-mitigation audits details.
type Verification struct {
	ID                string    `json:"id" db:"id"`
	HazardID          string    `json:"hazard_id" db:"hazard_id"`
	VerifierID        string    `json:"verifier_id" db:"verifier_id"`
	VerifiedDate      time.Time `json:"verified_date" db:"verified_date"`
	ResidualRiskScore int       `json:"residual_risk_score" db:"residual_risk_score"`
	Comments          string    `json:"comments,omitempty" db:"comments"`
}

// Validate checks domain invariants.
func (v *Verification) Validate() error {
	if v.HazardID == "" {
		return errors.New("hazard ID reference is required")
	}
	if v.VerifierID == "" {
		return errors.New("verifier ID reference is required")
	}
	if v.ResidualRiskScore < 1 || v.ResidualRiskScore > 25 {
		return errors.New("residual risk score must fall inside a 1-25 range")
	}
	return nil
}
