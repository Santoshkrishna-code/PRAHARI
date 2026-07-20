package policy

import (
	"time"

	"prahari/services/barrier/internal/domain/barrier"
)

// CalculateBarrierHealth computes overall health score based on PFD performance and proof testing.
func CalculateBarrierHealth(b *barrier.Barrier, lastTestPassed bool, activeImpairments int) float64 {
	score := 100.0
	if !lastTestPassed {
		score -= 40.0
	}
	if activeImpairments > 0 {
		score -= float64(activeImpairments) * 25.0
	}
	if b.NextProofTestDue != nil && b.NextProofTestDue.Before(time.Now()) {
		score -= 20.0
	}
	if score < 0 {
		score = 0
	}
	return score
}

// IsProofTestOverdue checks if proof testing deadline has passed.
func IsProofTestOverdue(b *barrier.Barrier) bool {
	if b.NextProofTestDue == nil {
		return false
	}
	return b.NextProofTestDue.Before(time.Now())
}
