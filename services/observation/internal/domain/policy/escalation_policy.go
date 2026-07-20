package policy

import (
	"errors"

	"prahari/services/observation/internal/domain/observation"
)

// VerifyEscalationParameters checks if critical unsafe acts trigger hazard logs.
func VerifyEscalationParameters(o *observation.Observation) error {
	if o.ObservationType == "CRITICAL_UNSAFE_ACT" && o.StatusCode == "REVIEWED" {
		return errors.New("critical unsafe acts require escalation to the Hazard Service as unsafe conditions")
	}
	return nil
}
