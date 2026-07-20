package search

import "time"

// Criteria defines multi-dimensional search parameters for PHA studies.
type Criteria struct {
	OrganizationID string     `json:"organization_id,omitempty"`
	PlantID        string     `json:"plant_id,omitempty"`
	UnitID         string     `json:"unit_id,omitempty"`
	ProcessNodeID  string     `json:"process_node_id,omitempty"`
	StudyType      string     `json:"study_type,omitempty"`
	HazardCategory string     `json:"hazard_category,omitempty"`
	RiskLevel      string     `json:"risk_level,omitempty"`
	RecStatus      string     `json:"rec_status,omitempty"`
	LeaderID       string     `json:"leader_id,omitempty"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
