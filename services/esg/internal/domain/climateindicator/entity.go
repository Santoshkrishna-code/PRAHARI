package climateindicator

import (
	"errors"
	"time"
)

// Indicator tracks environmental change risks.
type Indicator struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	RiskCategory   string    `json:"risk_category" db:"risk_category"` // "FLOOD", "WATER_STRESS", "CARBON_TAX"
	ImpactScore    int       `json:"impact_score" db:"impact_score"`   // 1 to 5 scale
	Probability    int       `json:"probability" db:"probability"`     // 1 to 5 scale
	FinancialRisk  float64   `json:"financial_risk" db:"financial_risk"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks indicators.
func (c *Indicator) Validate() error {
	if c.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}
