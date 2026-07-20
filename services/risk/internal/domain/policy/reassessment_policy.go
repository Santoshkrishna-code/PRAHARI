package policy

import (
	"errors"

	"prahari/services/risk/internal/domain/risk"
)

// CheckReassessmentParameters checks if critical incidents force safety updates.
func CheckReassessmentParameters(r *risk.Risk, incidentCount int) error {
	if incidentCount > 0 && r.StatusCode == "ACTIVE" {
		return errors.New("active risk zones with recorded incidents require urgent reassessment workflows")
	}
	return nil
}
