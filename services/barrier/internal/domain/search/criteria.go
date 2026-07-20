package search

import "time"

// Criteria defines multi-dimensional search parameters for safety barriers.
type Criteria struct {
	OrganizationID string     `json:"organization_id,omitempty"`
	PlantID        string     `json:"plant_id,omitempty"`
	UnitID         string     `json:"unit_id,omitempty"`
	BarrierType    string     `json:"barrier_type,omitempty"`
	AssetID        string     `json:"asset_id,omitempty"`
	SILLevel       string     `json:"sil_level,omitempty"`
	IsIPL          *bool      `json:"is_ipl,omitempty"`
	Status         string     `json:"status,omitempty"`
	IsImpaired     *bool      `json:"is_impaired,omitempty"`
	IsBypassed     *bool      `json:"is_bypassed,omitempty"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
