package lessonlearned

import "time"

// Lesson represents a systemic process safety lesson learned from an emergency.
type Lesson struct {
	ID          string    `json:"id"`
	ReviewID    string    `json:"review_id"`
	Category    string    `json:"category"` // EQUIPMENT, PROCEDURAL, TRAINING, COMMUNICATION
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ActionNeeded string   `json:"action_needed"`
	AssignedTo  string    `json:"assigned_to"`
	Status      string    `json:"status"` // OPEN, RESOLVED, VERIFIED
	CreatedAt   time.Time `json:"created_at"`
}
