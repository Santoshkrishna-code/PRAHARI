package policy

import (
	"errors"

	"prahari/services/maintenance/internal/domain/maintenance"
)

// ValidateMaintenanceCost checks approval thresholds.
func ValidateMaintenanceCost(m *maintenance.Maintenance) error {
	if m.TotalEstimatedCost > 50000.0 && m.StatusCode == "PLANNED" {
		return errors.New("maintenance work order estimated costs exceed $50k and require workflow approvals signoff")
	}
	return nil
}
