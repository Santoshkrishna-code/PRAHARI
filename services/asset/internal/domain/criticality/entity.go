package criticality

import (
	"errors"

	"prahari/services/asset/internal/domain/asset"
)

// Engine calculates risk severity.
type Engine struct {
	BusinessImpact      int `json:"business_impact"`      // 1 - 5
	SafetyImpact        int `json:"safety_impact"`        // 1 - 5
	EnvironmentalImpact int `json:"environmental_impact"` // 1 - 5
	ProductionImpact    int `json:"production_impact"`    // 1 - 5
}

// Validate checks limits.
func (e *Engine) Validate() error {
	if e.BusinessImpact < 1 || e.BusinessImpact > 5 {
		return errors.New("business impact must be between 1 and 5")
	}
	if e.SafetyImpact < 1 || e.SafetyImpact > 5 {
		return errors.New("safety impact must be between 1 and 5")
	}
	return nil
}

// CalculateCriticality calculates risk priority code levels.
func (e *Engine) CalculateCriticality() asset.CriticalityLevel {
	score := (float64(e.BusinessImpact) * 0.3) + (float64(e.SafetyImpact) * 0.3) + (float64(e.EnvironmentalImpact) * 0.2) + (float64(e.ProductionImpact) * 0.2)
	switch {
	case score >= 4.0:
		return asset.CritCritical
	case score >= 3.0:
		return asset.CritHigh
	case score >= 2.0:
		return asset.CritMedium
	default:
		return asset.CritLow
	}
}
