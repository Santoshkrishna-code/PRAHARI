package compliance

import (
	"prahari/services/chemical/internal/domain/chemical"
	"prahari/services/chemical/internal/domain/policy"
)

// ValidateSEVESO checks if a chemical inventory amount exceeds Seveso III threshold limits.
func ValidateSEVESO(c *chemical.Chemical, qty float64) (bool, string) {
	if c == nil {
		return false, "Chemical definition not found"
	}
	if policy.ExceedsSEVESOThreshold(c, qty) {
		return false, "Quantity exceeds SEVESO III threshold limits"
	}
	return true, "Complies with SEVESO III thresholds"
}
