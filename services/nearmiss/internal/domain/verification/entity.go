package verification

import (
	"errors"
	"time"
)

// Verification logs post-mitigation safety verifications.
type Verification struct {
	ID           string    `json:"id" db:"id"`
	NearMissID   string    `json:"near_miss_id" db:"near_miss_id"`
	VerifierID   string    `json:"verifier_id" db:"verifier_id"`
	VerifiedDate time.Time `json:"verified_date" db:"verified_date"`
	IsPassed     bool      `json:"is_passed" db:"is_passed"`
	Comments     string    `json:"comments,omitempty" db:"comments"`
}

// Validate checks domain invariants.
func (v *Verification) Validate() error {
	if v.NearMissID == "" {
		return errors.New("near miss ID reference is required")
	}
	if v.VerifierID == "" {
		return errors.New("verifier ID is required")
	}
	return nil
}
