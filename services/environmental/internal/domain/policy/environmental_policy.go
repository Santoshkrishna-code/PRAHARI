package policy

import (
	"errors"

	"prahari/services/environmental/internal/domain/emission"
	"prahari/services/environmental/internal/domain/waterquality"
)

// EvaluateEmissionLimits checks if gas stack emissions violate permit regulations.
func EvaluateEmissionLimits(e *emission.Emission) error {
	if e.ReleaseRate > e.LimitThreshold {
		return errors.New("atmospheric release rate violates statutory permit limit threshold")
	}
	return nil
}

// EvaluateWaterQuality evaluates pH and turbidity ranges.
func EvaluateWaterQuality(w *waterquality.WaterQuality) error {
	if w.PH < 6.5 || w.PH > 8.5 {
		return errors.New("effluent water pH violates discharge permit standard (must be 6.5 - 8.5)")
	}
	if w.TurbidityNTU > 5.0 {
		return errors.New("water turbidity level exceeds threshold limits")
	}
	return nil
}
