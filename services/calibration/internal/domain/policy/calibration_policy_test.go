package policy

import (
	"testing"
	"time"

	"prahari/services/calibration/internal/domain/calibrationschedule"
	"prahari/services/calibration/internal/domain/referencestandard"
	"prahari/services/calibration/internal/domain/tolerancerule"
)

func TestIsCalibrationOverdue(t *testing.T) {
	past := time.Now().Add(-24 * time.Hour)
	future := time.Now().Add(24 * time.Hour)

	schedPast := &calibrationschedule.Schedule{ScheduledFor: past}
	schedFuture := &calibrationschedule.Schedule{ScheduledFor: future}

	if !IsCalibrationOverdue(schedPast) {
		t.Errorf("Expected past schedule without completion to be overdue")
	}
	if IsCalibrationOverdue(schedFuture) {
		t.Errorf("Expected future schedule to NOT be overdue")
	}
}

func TestVerifyTolerance(t *testing.T) {
	rule := &tolerancerule.Rule{
		MinLimit: -0.5,
		MaxLimit: 0.5,
	}

	// 100 nominal, 100.2 measured -> 0.2% error, within ±0.5%
	if !VerifyTolerance(100.0, 100.2, rule) {
		t.Errorf("Expected 100.2 to be within tolerance limit")
	}

	// 100 nominal, 101 measured -> 1% error, exceeds limit
	if VerifyTolerance(100.0, 101.0, rule) {
		t.Errorf("Expected 101.0 to exceed tolerance limit")
	}
}

func TestIsReferenceStandardValid(t *testing.T) {
	future := time.Now().Add(24 * time.Hour)
	past := time.Now().Add(-24 * time.Hour)

	stdOk := &referencestandard.Standard{ExpiryDate: future}
	stdExpired := &referencestandard.Standard{ExpiryDate: past}

	if !IsReferenceStandardValid(stdOk) {
		t.Errorf("Expected standard with future expiry date to be valid")
	}
	if IsReferenceStandardValid(stdExpired) {
		t.Errorf("Expected expired standard to be invalid")
	}
}
