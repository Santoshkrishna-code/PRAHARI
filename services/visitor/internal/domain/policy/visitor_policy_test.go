package policy

import (
	"testing"
	"time"

	"prahari/services/visitor/internal/domain/blacklist"
	"prahari/services/visitor/internal/domain/induction"
)

func TestIsVisitorBlacklisted(t *testing.T) {
	entries := []blacklist.Entry{
		{IDNumber: "ID-12345", Reason: "Security breach"},
	}

	if !IsVisitorBlacklisted("ID-12345", entries) {
		t.Errorf("Expected visitor ID-12345 to be blacklisted")
	}
	if IsVisitorBlacklisted("ID-99999", entries) {
		t.Errorf("Expected visitor ID-99999 to NOT be blacklisted")
	}
}

func TestIsInductionValid(t *testing.T) {
	past := time.Now().Add(-24 * time.Hour)
	future := time.Now().Add(24 * time.Hour)

	verPast := &induction.Verification{ExpiresAt: past}
	verFuture := &induction.Verification{ExpiresAt: future}

	if IsInductionValid(verPast) {
		t.Errorf("Expected expired induction to be invalid")
	}
	if !IsInductionValid(verFuture) {
		t.Errorf("Expected non-expired induction to be valid")
	}
	if IsInductionValid(nil) {
		t.Errorf("Expected nil induction to be invalid")
	}
}

func TestCalculateGateAccessValidity(t *testing.T) {
	start := time.Now()
	validity := CalculateGateAccessValidity(start)

	expected := start.Add(8 * time.Hour)
	if validity.Sub(expected) > time.Second {
		t.Errorf("Expected gate validity to be 8 hours after start time")
	}
}
