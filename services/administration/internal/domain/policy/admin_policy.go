package policy

import (
	"time"

	"prahari/services/administration/internal/domain/featureflag"
	"prahari/services/administration/internal/domain/license"
)

// IsLicenseExpired checks if the tenant's license validity has passed.
func IsLicenseExpired(lic *license.License) bool {
	if lic == nil {
		return true
	}
	return time.Now().After(lic.ExpiresAt)
}

// IsPlantLimitReached checks if the tenant has already reached the licensed maximum plant count.
func IsPlantLimitReached(lic *license.License, currentPlantCount int) bool {
	if lic == nil {
		return true
	}
	return currentPlantCount >= lic.MaxPlants
}

// IsFeatureEnabled checks if a specific feature flag is turned on.
func IsFeatureEnabled(flag *featureflag.Flag) bool {
	if flag == nil {
		return false
	}
	return flag.Enabled
}
