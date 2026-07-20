package policy

import (
	"errors"

	"prahari/services/nearmiss/internal/domain/nearmiss"
)

// ValidateNearMissEscalation checks if serious potential level requires incident logs creation.
func ValidateNearMissEscalation(nm *nearmiss.NearMiss) error {
	if nm.SeverityLevel == "SERIOUS" && nm.StatusCode == "REPORTED" {
		return errors.New("serious near misses require escalation to the Incident Service for reactive review")
	}
	return nil
}
