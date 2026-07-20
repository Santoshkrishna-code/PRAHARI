package carboninventory

import (
	"errors"
	"time"
)

// Inventory structures the overall Scope emission inventory calculation runs.
type Inventory struct {
	ID             string    `json:"id" db:"id"`
	BusinessUnitID string    `json:"business_unit_id" db:"business_unit_id"`
	PeriodStart    time.Time `json:"period_start" db:"period_start"`
	PeriodEnd      time.Time `json:"period_end" db:"period_end"`
	Scope1Co2Kg    float64   `json:"scope_1_co2_kg" db:"scope_1_co2_kg"` // Direct emissions
	Scope2Co2Kg    float64   `json:"scope_2_co2_kg" db:"scope_2_co2_kg"` // Indirect (electricity) emissions
	Scope3Co2Kg    float64   `json:"scope_3_co2_kg" db:"scope_3_co2_kg"` // Supply chain/travel emissions
	TotalCo2Kg     float64   `json:"total_co2_kg" db:"total_co2_kg"`
	IsCalculated   bool      `json:"is_calculated" db:"is_calculated"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks limits.
func (i *Inventory) Validate() error {
	if i.BusinessUnitID == "" {
		return errors.New("business unit reference is required")
	}
	if i.PeriodEnd.Before(i.PeriodStart) {
		return errors.New("period end must be after period start")
	}
	return nil
}

// CalculateTotal calculates overall metric.
func (i *Inventory) CalculateTotal() {
	i.TotalCo2Kg = i.Scope1Co2Kg + i.Scope2Co2Kg + i.Scope3Co2Kg
	i.IsCalculated = true
}
