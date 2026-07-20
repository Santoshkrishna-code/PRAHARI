package wastewater

import (
	"errors"
	"time"
)

// Wastewater records effluent chemical discharge parameters.
type Wastewater struct {
	ID             string    `json:"id" db:"id"`
	OutfallID      string    `json:"outfall_id" db:"outfall_id"`
	BOD            float64   `json:"bod" db:"bod"` // Biological Oxygen Demand mg/L
	COD            float64   `json:"cod" db:"cod"` // Chemical Oxygen Demand mg/L
	TSS            float64   `json:"tss" db:"tss"` // Total Suspended Solids mg/L
	OilAndGrease   float64   `json:"oil_and_grease" db:"oil_and_grease"`
	FlowRateM3     float64   `json:"flow_rate_m3" db:"flow_rate_m3"` // Daily flow volume
	IsCompliant    bool      `json:"is_compliant" db:"is_compliant"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks wastewater fields.
func (w *Wastewater) Validate() error {
	if w.OutfallID == "" {
		return errors.New("outfall ID is required")
	}
	return nil
}
