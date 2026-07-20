package policy

import (
	"time"

	"prahari/services/visitor/internal/domain/blacklist"
	"prahari/services/visitor/internal/domain/induction"
)

// IsVisitorBlacklisted checks if the visitor details match a blacklist entry.
func IsVisitorBlacklisted(idNumber string, entries []blacklist.Entry) bool {
	for _, entry := range entries {
		if entry.IDNumber == idNumber {
			return true
		}
	}
	return false
}

// IsInductionValid checks if the visitor has a completed, non-expired safety induction.
func IsInductionValid(ver *induction.Verification) bool {
	if ver == nil {
		return false
	}
	return time.Now().Before(ver.ExpiresAt)
}

// CalculateGateAccessValidity computes default time limit for gate pass access (e.g. 8 hours).
func CalculateGateAccessValidity(from time.Time) time.Time {
	return from.Add(8 * time.Hour)
}
