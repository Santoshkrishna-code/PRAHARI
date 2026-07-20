package policy

import (
	"errors"

	"prahari/services/hazard/internal/domain/hazard"
)

// CheckEscalationRequirement checks if critical risk level triggers auto incident log.
func CheckEscalationRequirement(h *hazard.Hazard) error {
	if h.InitialRiskScore >= 20 && h.StatusCode == "RISK_ASSESSMENT" {
		return errors.New("critical initial risk scores require auto-escalation to the Incident Service for reactive review log")
	}
	return nil
}
