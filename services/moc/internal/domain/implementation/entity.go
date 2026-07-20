package implementation

import "time"

// Plan tracks physical / procedural execution of an approved change.
type Plan struct {
	ID              string     `json:"id"`
	ChangeRequestID string     `json:"change_request_id"`
	WorkOrderID     string     `json:"work_order_id,omitempty"` // Maintenance work order
	PermitID        string     `json:"permit_id,omitempty"`     // Permit-to-Work ID
	ImplementedBy   string     `json:"implemented_by"`
	Status          string     `json:"status"` // IN_PROGRESS, COMPLETED, FAILED
	StartDate       time.Time  `json:"start_date"`
	CompletedDate   *time.Time `json:"completed_date,omitempty"`
	Notes           string     `json:"notes"`
	CreatedAt       time.Time  `json:"created_at"`
}
