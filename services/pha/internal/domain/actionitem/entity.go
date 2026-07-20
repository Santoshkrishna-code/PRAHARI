package actionitem

import "time"

// ActionItem represents a concrete tracking action derived from a PHA recommendation.
type ActionItem struct {
	ID               string     `json:"id"`
	RecommendationID string     `json:"recommendation_id"`
	ActionTitle      string     `json:"action_title"`
	AssigneeID       string     `json:"assignee_id"`
	WorkOrderID      string     `json:"work_order_id,omitempty"` // Maintenance work order ID
	Status           string     `json:"status"`         // OPEN, IN_PROGRESS, COMPLETED, VERIFIED
	DueDate          time.Time  `json:"due_date"`
	CompletedAt      *time.Time `json:"completed_at,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
}
