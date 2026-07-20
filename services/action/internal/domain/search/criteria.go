package search

import "time"

// Criteria defines multi-dimensional search parameters for CAPAs, Actions, and Findings.
type Criteria struct {
	PlantID      string     `json:"plant_id,omitempty"`
	SourceModule string     `json:"source_module,omitempty"`
	Status       string     `json:"status,omitempty"`
	ActionType   string     `json:"action_type,omitempty"`
	AssignedTo   string     `json:"assigned_to,omitempty"`
	OverdueOnly  bool       `json:"overdue_only,omitempty"`
	DueDateAfter *time.Time `json:"due_date_after,omitempty"`
	Query        string     `json:"query,omitempty"`
	Limit        int        `json:"limit,omitempty"`
	Offset       int        `json:"offset,omitempty"`
}
