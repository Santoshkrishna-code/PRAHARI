package policy_test

import (
	"testing"
	"time"

	"prahari/services/administration/internal/domain/featureflag"
	"prahari/services/administration/internal/domain/license"
	"prahari/services/administration/internal/domain/policy"
)

func TestIsLicenseExpired(t *testing.T) {
	activeLic := &license.License{
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	expiredLic := &license.License{
		ExpiresAt: time.Now().Add(-24 * time.Hour),
	}

	if policy.IsLicenseExpired(activeLic) {
		t.Error("expected active license not to be expired")
	}

	if !policy.IsLicenseExpired(expiredLic) {
		t.Error("expected expired license to be flagged as expired")
	}
}

func TestIsPlantLimitReached(t *testing.T) {
	lic := &license.License{MaxPlants: 5}

	if !policy.IsPlantLimitReached(lic, 5) {
		t.Error("expected limit reached with 5 plants out of 5 maximum")
	}

	if !policy.IsPlantLimitReached(lic, 6) {
		t.Error("expected limit reached with 6 plants out of 5 maximum")
	}

	if policy.IsPlantLimitReached(lic, 4) {
		t.Error("expected limit not reached with 4 plants out of 5 maximum")
	}

}

func TestIsFeatureEnabled(t *testing.T) {
	enabledFlag := &featureflag.Flag{Enabled: true}
	disabledFlag := &featureflag.Flag{Enabled: false}

	if !policy.IsFeatureEnabled(enabledFlag) {
		t.Error("expected enabled flag to return true")
	}

	if policy.IsFeatureEnabled(disabledFlag) {
		t.Error("expected disabled flag to return false")
	}
}
