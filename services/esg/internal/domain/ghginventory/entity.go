package ghginventory

import (
	"errors"
	"time"
)

// GHGRecord tracks separate gas metrics.
type GHGRecord struct {
	ID          string    `json:"id" db:"id"`
	InventoryID string    `json:"inventory_id" db:"inventory_id"`
	GasType     string    `json:"gas_type" db:"gas_type"` // "CO2", "CH4", "N2O", "HFCs"
	VolumeTonnes float64  `json:"volume_tonnes" db:"volume_tonnes"`
	GWPFactor   float64   `json:"gwp_factor" db:"gwp_factor"` // Global Warming Potential factor
	Co2EqTonnes float64   `json:"co2_eq_tonnes" db:"co2_eq_tonnes"`
	RecordedAt  time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks fields.
func (g *GHGRecord) Validate() error {
	if g.InventoryID == "" {
		return errors.New("inventory reference is required")
	}
	if g.GasType == "" {
		return errors.New("gas type classification is required")
	}
	return nil
}

// CalculateEquivalent runs multiplier check.
func (g *GHGRecord) CalculateEquivalent() {
	g.Co2EqTonnes = g.VolumeTonnes * g.GWPFactor
}
