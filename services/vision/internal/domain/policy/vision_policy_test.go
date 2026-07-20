package policy_test

import (
	"testing"

	"prahari/services/vision/internal/domain/policy"
)

func TestIsWithinZone(t *testing.T) {
	// Object centers at x=200, y=200. Zone is [100, 100, 300, 300]
	if !policy.IsWithinZone(150, 150, 100, 100, 100, 100, 300, 300) {
		t.Error("expected center of bounding box to fall within zone boundaries")
	}

	// Object centers at x=50, y=50
	if policy.IsWithinZone(0, 0, 100, 100, 100, 100, 300, 300) {
		t.Error("expected center of bounding box to fall outside zone boundaries")
	}
}

func TestIsDetectionValid(t *testing.T) {
	if !policy.IsDetectionValid(0.95, 0.8) {
		t.Error("expected confidence score of 0.95 to exceed threshold 0.8")
	}

	if policy.IsDetectionValid(0.6, 0.8) {
		t.Error("expected confidence score of 0.6 not to exceed threshold 0.8")
	}
}
