package policy

import (
	"time"

	"prahari/services/calibration/internal/domain/calibrationschedule"
	"prahari/services/calibration/internal/domain/referencestandard"
	"prahari/services/calibration/internal/domain/tolerancerule"
)


// IsCalibrationOverdue checks if the scheduled calibration date has passed without completion.
func IsCalibrationOverdue(sched *calibrationschedule.Schedule) bool {
	return sched.CompletedAt == nil && time.Now().After(sched.ScheduledFor)
}

// VerifyTolerance checks if the difference between nominal value and as-found value falls within min/max tolerance rule limits.
func VerifyTolerance(nominal, measured float64, rule *tolerancerule.Rule) bool {
	if rule == nil {
		return true
	}
	errorVal := measured - nominal
	percentError := (errorVal / nominal) * 100.0
	return percentError >= rule.MinLimit && percentError <= rule.MaxLimit
}

// IsReferenceStandardValid checks if reference standards are within valid calibration certification periods.
func IsReferenceStandardValid(std *referencestandard.Standard) bool {
	if std == nil {
		return false
	}
	return time.Now().Before(std.ExpiryDate)
}
