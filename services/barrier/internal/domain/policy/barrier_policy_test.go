package policy

import (
	"testing"
	"time"

	"prahari/services/barrier/internal/domain/barrier"
)

func TestCalculateBarrierHealth(t *testing.T) {
	past := time.Now().Add(-24 * time.Hour)
	b := &barrier.Barrier{
		NextProofTestDue: &past,
	}

	score := CalculateBarrierHealth(b, false, 1)
	// 100 - 40 (failed test) - 25 (1 impairment) - 20 (overdue) = 15.0
	if score != 15.0 {
		t.Errorf("CalculateBarrierHealth() got score=%.2f; want 15.00", score)
	}
}

func TestIsProofTestOverdue(t *testing.T) {
	past := time.Now().Add(-24 * time.Hour)
	bOverdue := &barrier.Barrier{NextProofTestDue: &past}
	if !IsProofTestOverdue(bOverdue) {
		t.Errorf("Expected barrier with past proof test date to be overdue")
	}
}
