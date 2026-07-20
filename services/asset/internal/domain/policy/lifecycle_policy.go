package policy

import (
	"errors"

	"prahari/services/asset/internal/domain/asset"
)

// ValidateLifecyclePermitState checks state boundaries.
func ValidateLifecyclePermitState(a *asset.Asset) error {
	if a.LifecycleStatus == "DECOMMISSIONED" || a.LifecycleStatus == "DISPOSED" {
		return errors.New("decommissioned assets cannot accept operational permit requests")
	}
	return nil
}
