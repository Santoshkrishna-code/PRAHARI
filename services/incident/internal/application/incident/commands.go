package incident

import (
	"time"
)

// CreateIncidentCommand carries the data required to register a new incident.
type CreateIncidentCommand struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Type           string    `json:"type"`
	CategoryID     string    `json:"category_id"`
	SeverityLevel  string    `json:"severity_level"`
	PriorityLevel  string    `json:"priority_level"`
	ReporterID     string    `json:"reporter_id"`
	DepartmentID   string    `json:"department_id"`
	LocationID     string    `json:"location_id"`
	LocationDetail string    `json:"location_detail,omitempty"`
	OccurredAt     time.Time `json:"occurred_at"`
}

// UpdateIncidentCommand carries the data required to modify an existing incident.
type UpdateIncidentCommand struct {
	Title          string `json:"title,omitempty"`
	Description    string `json:"description,omitempty"`
	CategoryID     string `json:"category_id,omitempty"`
	SeverityLevel  string `json:"severity_level,omitempty"`
	PriorityLevel  string `json:"priority_level,omitempty"`
	LocationID     string `json:"location_id,omitempty"`
	LocationDetail string `json:"location_detail,omitempty"`
}

// TransitionStatusCommand carries the data required to change an incident's status.
type TransitionStatusCommand struct {
	IncidentID string `json:"incident_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}
