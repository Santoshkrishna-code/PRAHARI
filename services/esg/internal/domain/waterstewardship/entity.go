package waterstewardship

import (
	"errors"
	"time"
)

// StewardshipKPI tracks water usage metrics.
type StewardshipKPI struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	WithdrawalM3   float64   `json:"withdrawal_m3" db:"withdrawal_m3"` // Total water taken in
	ConsumptionM3  float64   `json:"consumption_m3" db:"consumption_m3"` // Withdrawal - Discharge
	RecycledM3     float64   `json:"recycled_m3" db:"recycled_m3"`
	RecycleRatio   float64   `json:"recycle_ratio" db:"recycle_ratio"` // Recycled / Withdrawal
	RecordedPeriod string    `json:"recorded_period" db:"recorded_period"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks fields.
func (w *StewardshipKPI) Validate() error {
	if w.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}

// EvaluateRatio calculates rates.
func (w *StewardshipKPI) EvaluateRatio() {
	if w.WithdrawalM3 > 0.0 {
		w.RecycleRatio = w.RecycledM3 / w.WithdrawalM3
	}
}
