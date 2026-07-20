package policy

import (
	"prahari/services/bcm/internal/domain/businessimpactanalysis"
	"prahari/services/bcm/internal/domain/criticalprocess"
)

// DeterminePriorityTier evaluates Maximum Tolerable Downtime (MTD) to assign process priority.
func DeterminePriorityTier(mtdHrs float64) string {
	switch {
	case mtdHrs <= 4.0:
		return "TIER_1_CRITICAL"
	case mtdHrs <= 24.0:
		return "TIER_2_IMPORTANT"
	default:
		return "TIER_3_NORMAL"
	}
}

// IsRTOCompliant checks if actual recovery time met target Recovery Time Objective (RTO).
func IsRTOCompliant(actualRTOHrs, targetRTOHrs float64) bool {
	return actualRTOHrs <= targetRTOHrs
}

// CalculateResilienceScore computes overall BCM score for a critical process.
func CalculateResilienceScore(cp *criticalprocess.Process, bia *businessimpactanalysis.Analysis, drTested bool) float64 {
	score := 100.0
	if bia == nil {
		score -= 40.0
	}
	if !drTested {
		score -= 30.0
	}
	if score < 0 {
		score = 0
	}
	return score
}
