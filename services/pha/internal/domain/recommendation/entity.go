package recommendation

import "time"

// Recommendation represents a process safety improvement recommendation resulting from PHA.
type Recommendation struct {
	ID           string    `json:"id"`
	StudyID      string    `json:"study_id"`
	ScenarioID   string    `json:"scenario_id,omitempty"`
	RecNumber    string    `json:"rec_number"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Priority     string    `json:"priority"` // LOW, MEDIUM, HIGH, CRITICAL
	TargetSIL    string    `json:"target_sil,omitempty"` // SIL-1, SIL-2, SIL-3 (if LOPA driven)
	Status       string    `json:"status"`   // OPEN, IN_PROGRESS, RESOLVED, CLOSED
	AssignedTo   string    `json:"assigned_to"`
	TargetDate   time.Time `json:"target_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
