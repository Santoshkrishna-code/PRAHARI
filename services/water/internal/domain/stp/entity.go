package stp

import "time"

// STP represents a Sewage Treatment Plant facility.
type STP struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	FacilityName    string    `json:"facility_name"`
	TechnologyType  string    `json:"technology_type"`
	DesignCapacityKLD float64 `json:"design_capacity_kld"`
	InfluentKLD     float64   `json:"influent_kld"`
	EffluentKLD     float64   `json:"effluent_kld"`
	BODInfluent     float64   `json:"bod_influent_mg_l"`
	BODEffluent     float64   `json:"bod_effluent_mg_l"`
	CODInfluent     float64   `json:"cod_influent_mg_l"`
	CODEffluent     float64   `json:"cod_effluent_mg_l"`
	TSSEffluent     float64   `json:"tss_effluent_mg_l"`
	AssetID         string    `json:"asset_id,omitempty"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
