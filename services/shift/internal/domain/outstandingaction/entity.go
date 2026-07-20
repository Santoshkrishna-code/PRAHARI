package outstandingaction

import "time"

// Action represents an outstanding task to be resolved across shift cycles.
type Action struct {
	ID          string     `json:"id"`
	ShiftID     string     `json:"shift_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	AssignedTo  string     `json:"assigned_to"`
	DueDate     time.Time  `json:"due_date"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Status      string     `json:"status"` // OPEN, IN_PROGRESS, COMPLETED, OVERDUE
}
