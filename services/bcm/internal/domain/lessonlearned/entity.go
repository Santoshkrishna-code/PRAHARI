package lessonlearned

import "time"

// Lesson represents a continuous improvement lesson learned from BCM activation or exercise.
type Lesson struct {
	ID          string    `json:"id"`
	ReviewID    string    `json:"review_id"`
	Category    string    `json:"category"` // STRATEGY, DR_TECHNICAL, SUPPLIER, GOVERNANCE
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ActionNeeded string   `json:"action_needed"`
	AssignedTo  string    `json:"assigned_to"`
	Status      string    `json:"status"` // OPEN, RESOLVED, VERIFIED
	CreatedAt   time.Time `json:"created_at"`
}
