package policy

import (
	"errors"

	"prahari/services/asset/internal/domain/asset"
)

// ValidateCommissionPreconditions checks safety inspection signoffs.
func ValidateCommissionPreconditions(a *asset.Asset, hasPSSRPassed bool) error {
	if a.CriticalityCode == asset.CritCritical && !hasPSSRPassed {
		return errors.New("critical assets must pass pre-startup safety reviews before commissioning")
	}
	return nil
}
