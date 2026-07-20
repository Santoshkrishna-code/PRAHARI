package compliance

import (
	"prahari/services/chemical/internal/domain/ghsclassification"
)

// ValidateGHS ensures that the safety data contains required hazard details.
func ValidateGHS(g *ghsclassification.GHS) (bool, string) {
	if g == nil {
		return false, "GHS classification is missing"
	}
	if g.SignalWord == "" {
		return false, "Signal word is missing"
	}
	return true, "GHS data valid"
}
