package energyconsumption

import (
	"errors"
	"time"
)

// Consumption tracks aggregated load parameters.
type Consumption struct {
	ID             string    `json:"id" db:"id"`
	MeterID        string    `json:"meter_id" db:"meter_id"`
	PeriodStart    time.Time `json:"period_start" db:"period_start"`
	PeriodEnd      time.Time `json:"period_end" db:"period_end"`
	ConsumptionKWh float64   `json:"consumption_kwh" db:"consumption_kwh"`
	PeakDemandKW   float64   `json:"peak_demand_kw" db:"peak_demand_kw"`
	CarbonEmittedKg float64  `json:"carbon_emitted_kg" db:"carbon_emitted_kg"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks limits.
func (c *Consumption) Validate() error {
	if c.MeterID == "" {
		return errors.New("meter reference ID is required")
	}
	if c.PeriodEnd.Before(c.PeriodStart) {
		return errors.New("period end must be after period start")
	}
	return nil
}
