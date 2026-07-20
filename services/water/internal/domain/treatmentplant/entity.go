package treatmentplant

import "time"

// TreatmentPlant represents a water treatment facility (WTP).
type TreatmentPlant struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	FacilityName    string    `json:"facility_name"`
	TreatmentType   string    `json:"treatment_type"`
	DesignCapacityKLD float64 `json:"design_capacity_kld"`
	OperatingKLD    float64   `json:"operating_kld"`
	EfficiencyPct   float64   `json:"efficiency_pct"`
	AssetID         string    `json:"asset_id,omitempty"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
