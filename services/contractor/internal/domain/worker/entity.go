package worker

import (
	"errors"
)

// Worker represents an individual employee contractor crew member.
type Worker struct {
	ID               string `json:"id" db:"id"`
	ContractorID     string `json:"contractor_id" db:"contractor_id"`
	FirstName        string `json:"first_name" db:"first_name"`
	LastName         string `json:"last_name" db:"last_name"`
	PassportID       string `json:"passport_id" db:"passport_id"`
	OnboardingStatus string `json:"onboarding_status" db:"onboarding_status"` // Pending, Active, Blacklisted
}

// Validate checks domain invariants.
func (w *Worker) Validate() error {
	if w.ContractorID == "" {
		return errors.New("contractor reference is required")
	}
	if w.FirstName == "" || w.LastName == "" {
		return errors.New("first and last name are required")
	}
	return nil
}
