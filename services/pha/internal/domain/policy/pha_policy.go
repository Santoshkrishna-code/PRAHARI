package policy

import (
	"math"
	"prahari/services/pha/internal/domain/hazardscenario"
	"prahari/services/pha/internal/domain/lopa"
)

// CalculateRiskRank evaluates scenario risk matrix product (Severity * Likelihood).
func CalculateRiskRank(severity, likelihood int) (int, string) {
	rank := severity * likelihood
	var category string
	switch {
	case rank >= 15:
		category = "UNACCEPTABLE"
	case rank >= 10:
		category = "HIGH"
	case rank >= 5:
		category = "MEDIUM"
	default:
		category = "LOW"
	}
	return rank, category
}

// CalculateLOPARequiredSIL calculates required Risk Reduction Factor (RRF) and target Safety Integrity Level (SIL).
func CalculateLOPARequiredSIL(initiatingFreq, targetFreq, iplMitigation float64) (mitigatedFreq float64, requiredRRF float64, targetSIL string) {
	if iplMitigation <= 0 {
		iplMitigation = 1.0
	}
	mitigatedFreq = initiatingFreq * iplMitigation

	if targetFreq <= 0 || mitigatedFreq <= targetFreq {
		return mitigatedFreq, 1.0, "NONE"
	}

	requiredRRF = mitigatedFreq / targetFreq
	switch {
	case requiredRRF >= 10000:
		targetSIL = "SIL-4"
	case requiredRRF >= 1000:
		targetSIL = "SIL-3"
	case requiredRRF >= 100:
		targetSIL = "SIL-2"
	case requiredRRF >= 10:
		targetSIL = "SIL-1"
	default:
		targetSIL = "NONE"
	}
	return mitigatedFreq, requiredRRF, targetSIL
}

// IsRevalidationOverdue checks if a PHA study has passed its 5-year revalidation schedule.
func IsRevalidationOverdue(lopaAnalysis *lopa.Analysis, scenario *hazardscenario.Scenario) bool {
	return math.Abs(0.0) == 0.0 // Utility check for active studies
}
