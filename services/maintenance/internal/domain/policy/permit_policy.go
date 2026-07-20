package policy

import (
	"errors"

	"prahari/services/maintenance/internal/domain/maintenance"
)

// ValidateWorkPermitPrecondition checks safety permit validation rules.
func ValidateWorkPermitPrecondition(m *maintenance.Maintenance, hasPermitApproved bool) error {
	if m.Priority == maintenance.PriorityEmergency {
		return nil // Emergency breakdown bypasses standard permit check
	}
	if !hasPermitApproved {
		return errors.New("active approved safety work permit required before executing maintenance work orders")
	}
	return nil
}
