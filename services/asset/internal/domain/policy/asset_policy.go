package policy

import (
	"errors"

	"prahari/services/asset/internal/domain/asset"
)

// ValidateAssetMutation enforces immutable rules on Retired/Disposed assets.
func ValidateAssetMutation(a *asset.Asset) error {
	if a.LifecycleStatus == "DISPOSED" {
		return errors.New("cannot modify disposed equipment assets")
	}
	return nil
}
