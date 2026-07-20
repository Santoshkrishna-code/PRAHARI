package action

import "time"

// Action represents a single correction, prevention, or improvement task with a due date.
type Action struct {
	ID             string     `json:"id"`
	PlantID        string     `json:"plant_id"`
	SourceModule   string     `json:"source_module"` // INCIDENT, AUDIT, HAZARD, INSPECTION, PHA, etc.
	SourceRefID    string     `json:"source_ref_id"`  // ID of the source record
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	ActionType     string     `json:"action_type"` // CORRECTIVE, PREVENTIVE, IMPROVEMENT
	Status         string     `json:"status"`      // Created, Assigned, In Progress, Evidence Submitted, Effectiveness Review, Closed, Cancelled, Overdue, Rejected
	AssignedTo     string     `json:"assigned_to,omitempty"`
	DueDate        time.Time  `json:"due_date"`
	ClosedAt       *time.Time `json:"closed_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
