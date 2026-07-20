package etp

import "time"

// ETP represents an Effluent Treatment Plant facility.
type ETP struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	FacilityName    string    `json:"facility_name"`
	TechnologyType  string    `json:"technology_type"`
	DesignCapacityKLD float64 `json:"design_capacity_kld"`
	InfluentKLD     float64   `json:"influent_kld"`
	EffluentKLD     float64   `json:"effluent_kld"`
	PHInfluent      float64   `json:"ph_influent"`
	PHEffluent      float64   `json:"ph_effluent"`
	CODInfluent     float64   `json:"cod_influent_mg_l"`
	CODEffluent     float64   `json:"cod_effluent_mg_l"`
	TDSEffluent     float64   `json:"tds_effluent_mg_l"`
	OilGreaseEffl   float64   `json:"oil_grease_effluent_mg_l"`
	AssetID         string    `json:"asset_id,omitempty"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
