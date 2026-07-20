package policy

import (
	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/lock"
	"prahari/services/loto/internal/domain/verification"
)

// CanCommenceMaintenance checks if zero-energy has been verified so work can begin.
func CanCommenceMaintenance(cert *isolationcertificate.Certificate, verify *verification.ZeroEnergy) bool {
	if cert == nil || verify == nil {
		return false
	}
	return cert.Status == "ZERO_ENERGY_VERIFIED" && verify.TestPassed
}

// AllLocksRemoved checks if all assigned locks are available/removed.
func AllLocksRemoved(locks []*lock.Lock) bool {
	for _, l := range locks {
		if l.Status == "APPLIED" {
			return false
		}
	}
	return true
}
