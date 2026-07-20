package policy

import (
	"errors"
	"time"

	"prahari/services/contractor/internal/domain/contractor"
)

// ValidateContractorInsurance checks liability validation rules.
func ValidateContractorInsurance(c *contractor.Contractor) error {
	if time.Now().After(c.InsuranceExpiry) {
		return errors.New("contractor company insurance policies have expired")
	}
	return nil
}
