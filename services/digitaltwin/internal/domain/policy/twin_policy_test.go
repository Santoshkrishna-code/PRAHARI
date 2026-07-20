package policy_test

import (
	"testing"

	"prahari/services/digitaltwin/internal/domain/policy"
)

func TestValidateSimulationParameters(t *testing.T) {
	if !policy.ValidateSimulationParameters(`{"pressure": 12.0}`) {
		t.Error("expected valid parameters string to pass validation")
	}

	if policy.ValidateSimulationParameters("") {
		t.Error("expected empty parameters to fail validation")
	}

	// Create oversized string
	oversized := make([]byte, 60000)
	if policy.ValidateSimulationParameters(string(oversized)) {
		t.Error("expected oversized parameters to fail validation")
	}
}

func TestValidateTelemetryQuality(t *testing.T) {
	if !policy.ValidateTelemetryQuality("GOOD") {
		t.Error("expected GOOD quality status to be valid")
	}

	if !policy.ValidateTelemetryQuality("UNCERTAIN") {
		t.Error("expected UNCERTAIN quality status to be valid")
	}

	if policy.ValidateTelemetryQuality("BAD") {
		t.Error("expected BAD quality status to be invalid")
	}
}
