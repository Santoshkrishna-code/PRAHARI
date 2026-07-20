package policy

import (
	"errors"

	"prahari/services/contractor/internal/domain/contractor"
)

// ValidateWorkerInduction check site access safety orientation preconditions.
func ValidateWorkerInduction(c *contractor.Contractor, hasPassedSafetyInduction bool) error {
	if c.StatusCode == "ACTIVE" && !hasPassedSafetyInduction {
		return errors.New("contractor worker must pass site safety inductions prior to gate access")
	}
	return nil
}
