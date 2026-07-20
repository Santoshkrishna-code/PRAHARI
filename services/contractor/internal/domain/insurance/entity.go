package insurance

import (
	"errors"
	"time"
)

// Insurance validates contractor company insurance coverage.
type Insurance struct {
	ID          string    `json:"id" db:"id"`
	ContractorID string   `json:"contractor_id" db:"contractor_id"`
	PolicyNumber string    `json:"policy_number" db:"policy_number"`
	ExpiryDate  time.Time `json:"expiry_date" db:"expiry_date"`
	LimitAmount float64   `json:"limit_amount" db:"limit_amount"`
}

// Validate checks domain invariants.
func (i *Insurance) Validate() error {
	if i.ContractorID == "" {
		return errors.New("contractor ID is required")
	}
	if i.PolicyNumber == "" {
		return errors.New("policy number is required")
	}
	return nil
}

// IsValid checks validity windows.
func (i *Insurance) IsValid() bool {
	return time.Now().Before(i.ExpiryDate)
}
