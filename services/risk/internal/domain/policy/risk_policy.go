package policy

import (
	"errors"

	"prahari/services/risk/internal/domain/risk"
)

// CheckRiskEscalationRequirements checks if high inherent risks require dynamic controls evaluations.
func CheckRiskEscalationRequirements(r *risk.Risk) error {
	if r.InherentRiskScore >= 15 && r.StatusCode == "ASSESSMENT" {
		return errors.New("high inherent risks require mandatory bow-tie barriers analysis before review submissions")
	}
	return nil
}
