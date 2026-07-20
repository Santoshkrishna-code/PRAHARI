package desalination

import "time"

// Plant represents a desalination facility.
type Plant struct {
	ID               string    `json:"id"`
	PlantID          string    `json:"plant_id"`
	FacilityName     string    `json:"facility_name"`
	Technology       string    `json:"technology"`
	DesignCapacityKLD float64  `json:"design_capacity_kld"`
	OperatingKLD     float64   `json:"operating_kld"`
	RecoveryRatePct  float64   `json:"recovery_rate_pct"`
	EnergyPerKLKWh   float64   `json:"energy_per_kl_kwh"`
	BrineDischargeKLD float64  `json:"brine_discharge_kld"`
	AssetID          string    `json:"asset_id,omitempty"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
