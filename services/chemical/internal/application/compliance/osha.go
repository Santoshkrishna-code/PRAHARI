package compliance

import (
	"prahari/services/chemical/internal/domain/chemical"
	"prahari/services/chemical/internal/domain/policy"
)

// ValidateOSHAPSM checks if a chemical inventory amount exceeds OSHA PSM threshold limits.
func ValidateOSHAPSM(c *chemical.Chemical, qty float64) (bool, string) {
	if c == nil {
		return false, "Chemical definition not found"
	}
	if policy.ExceedsOSHAPSMThreshold(c, qty) {
		return false, "Quantity exceeds OSHA PSM threshold limit (TQ)"
	}
	return true, "Complies with OSHA PSM thresholds"
}
