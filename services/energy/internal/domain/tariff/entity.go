package tariff

import (
	"errors"
	"time"
)

// Structure maps electricity utility cost schemes.
type Structure struct {
	ID             string    `json:"id" db:"id"`
	ProviderName   string    `json:"provider_name" db:"provider_name"` // "State Utility Co"
	TariffName     string    `json:"tariff_name" db:"tariff_name"`     // "Industrial Peak Rate"
	PeakRateKWh    float64   `json:"peak_rate_kwh" db:"peak_rate_kwh"` // USD rate
	OffPeakRateKWh float64   `json:"off_peak_rate_kwh" db:"off_peak_rate_kwh"`
	DemandChargeKW float64   `json:"demand_charge_kw" db:"demand_charge_kw"`
	EffectiveFrom  time.Time `json:"effective_from" db:"effective_from"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks tariff.
func (s *Structure) Validate() error {
	if s.ProviderName == "" {
		return errors.New("provider name key is required")
	}
	if s.TariffName == "" {
		return errors.New("tariff plan name is required")
	}
	return nil
}
