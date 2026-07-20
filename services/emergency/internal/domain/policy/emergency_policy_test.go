package policy

import (
	"testing"

	"prahari/services/emergency/internal/domain/emergency"
	"prahari/services/emergency/internal/domain/evacuation"
)

func TestCalculateResponseEffectiveness(t *testing.T) {
	em := &emergency.Emergency{Severity: "TIER_3"}
	evac := &evacuation.Record{TotalPersonnel: 100, MissingCount: 2}

	score := CalculateResponseEffectiveness(em, evac)
	// 100 - (2/100 * 100 * 5) = 100 - 10 = 90.0
	if score != 90.0 {
		t.Errorf("CalculateResponseEffectiveness() got score=%.2f; want 90.00", score)
	}
}

func TestRequiresMutualAid(t *testing.T) {
	if !RequiresMutualAid("TIER_3") {
		t.Errorf("Expected TIER_3 severity to require mutual aid")
	}
	if RequiresMutualAid("TIER_1") {
		t.Errorf("Expected TIER_1 severity to NOT require mutual aid")
	}
}
