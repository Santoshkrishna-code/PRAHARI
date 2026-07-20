package policy

import (
	"time"

	"prahari/services/shift/internal/domain/checklist"
)

// CanInitiateHandover verifies if all shift start and operational checklists are signed off before handover initiation.
func CanInitiateHandover(items []checklist.Item) bool {
	for _, item := range items {
		if !item.Completed {
			return false
		}
	}
	return true
}

// CalculateShiftDuration computes total hours from actual start to end of shift.
func CalculateShiftDuration(start time.Time, end time.Time) float64 {
	return end.Sub(start).Hours()
}

// IsOvertimeRequired determines if shift duration exceeds standard shift duration limit (e.g. 8.0 hours).
func IsOvertimeRequired(durationHrs float64) bool {
	return durationHrs > 8.0
}
