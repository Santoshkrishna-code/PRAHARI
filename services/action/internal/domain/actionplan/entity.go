package actionplan

import "time"

// Plan outlines a scheduled timeline of implementation tasks/phases to achieve corrective results.
type Plan struct {
	ID          string    `json:"id"`
	CapaID      string    `json:"capa_id"`
	ApproverID  string    `json:"approver_id,omitempty"`
	ApprovedAt  *time.Time `json:"approved_at,omitempty"`
	TargetDate  time.Time `json:"target_date"`
	Description string    `json:"description"`
}
