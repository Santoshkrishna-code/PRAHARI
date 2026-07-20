package energycost

import (
	"errors"
	"time"
)

// CostRecord logs financial energy expenses.
type CostRecord struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	BillingPeriod  string    `json:"billing_period" db:"billing_period"` // "2026-07"
	TotalCostUSD   float64   `json:"total_cost_usd" db:"total_cost_usd"`
	ConsumptionKWh float64   `json:"consumption_kwh" db:"consumption_kwh"`
	AverageRateKWh float64   `json:"average_rate_kwh" db:"average_rate_kwh"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks fields.
func (c *CostRecord) Validate() error {
	if c.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}

// EvaluateAverage calculates rates.
func (c *CostRecord) EvaluateAverage() {
	if c.ConsumptionKWh > 0.0 {
		c.AverageRateKWh = c.TotalCostUSD / c.ConsumptionKWh
	}
}
