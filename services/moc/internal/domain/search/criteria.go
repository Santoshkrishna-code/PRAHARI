package search

import "time"

// Criteria defines multi-dimensional search parameters for MOC requests.
type Criteria struct {
	OrganizationID string     `json:"organization_id,omitempty"`
	PlantID        string     `json:"plant_id,omitempty"`
	DepartmentID   string     `json:"department_id,omitempty"`
	ChangeType     string     `json:"change_type,omitempty"`
	AssetID        string     `json:"asset_id,omitempty"`
	RiskLevel      string     `json:"risk_level,omitempty"`
	Status         string     `json:"status,omitempty"`
	RequesterID    string     `json:"requester_id,omitempty"`
	ApproverID     string     `json:"approver_id,omitempty"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
