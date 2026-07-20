package search

import "time"

// Criteria defines multi-dimensional search parameters for business continuity plans and BIA reports.
type Criteria struct {
	OrganizationID string     `json:"organization_id,omitempty"`
	PlantID        string     `json:"plant_id,omitempty"`
	BusinessUnit   string     `json:"business_unit,omitempty"`
	CriticalProcess string    `json:"critical_process,omitempty"`
	ContinuityPlan string     `json:"continuity_plan,omitempty"`
	Status         string     `json:"status,omitempty"`
	PriorityTier   string     `json:"priority_tier,omitempty"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
