package policy

import (
	"errors"

	"prahari/services/hazard/internal/domain/hazard"
)

// ValidateResidualRisk checks mitigation limits score rules.
func ValidateResidualRisk(h *hazard.Hazard) error {
	if h.ResidualRiskScore > 16 && h.StatusCode == "VERIFICATION" {
		return errors.New("residual risk score is still high/critical post-mitigation; additional safety control barriers are required")
	}
	return nil
}
