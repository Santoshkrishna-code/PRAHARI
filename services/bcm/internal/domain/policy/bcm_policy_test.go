package policy

import (
	"testing"
)

func TestDeterminePriorityTier(t *testing.T) {
	if DeterminePriorityTier(2.0) != "TIER_1_CRITICAL" {
		t.Errorf("DeterminePriorityTier(2.0) want TIER_1_CRITICAL")
	}
	if DeterminePriorityTier(12.0) != "TIER_2_IMPORTANT" {
		t.Errorf("DeterminePriorityTier(12.0) want TIER_2_IMPORTANT")
	}
	if DeterminePriorityTier(48.0) != "TIER_3_NORMAL" {
		t.Errorf("DeterminePriorityTier(48.0) want TIER_3_NORMAL")
	}
}

func TestIsRTOCompliant(t *testing.T) {
	if !IsRTOCompliant(3.5, 4.0) {
		t.Errorf("IsRTOCompliant(3.5, 4.0) want true")
	}
	if IsRTOCompliant(5.0, 4.0) {
		t.Errorf("IsRTOCompliant(5.0, 4.0) want false")
	}
}
