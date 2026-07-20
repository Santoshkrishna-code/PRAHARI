package policy

import (
	"prahari/services/emergency/internal/domain/emergency"
	"prahari/services/emergency/internal/domain/evacuation"
)

// CalculateResponseEffectiveness evaluates response time and evacuation completion efficiency.
func CalculateResponseEffectiveness(em *emergency.Emergency, evac *evacuation.Record) float64 {
	score := 100.0
	if evac != nil && evac.TotalPersonnel > 0 {
		missingPct := (float64(evac.MissingCount) / float64(evac.TotalPersonnel)) * 100.0
		score -= missingPct * 5.0
	}
	if score < 0 {
		score = 0
	}
	return score
}

// RequiresMutualAid checks if high-severity industrial emergencies require external mutual aid activation.
func RequiresMutualAid(severity string) bool {
	return severity == "TIER_3" || severity == "MAJOR"
}
