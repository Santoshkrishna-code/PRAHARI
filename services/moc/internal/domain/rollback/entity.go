package rollback

import "time"

// Plan outlines instructions to safely revert a failed or expired temporary change.
type Plan struct {
	ID              string     `json:"id"`
	ChangeRequestID string     `json:"change_request_id"`
	TriggerReason   string     `json:"trigger_reason"`
	ReversionSteps  string     `json:"reversion_steps"`
	ExecutedBy      string     `json:"executed_by,omitempty"`
	Status          string     `json:"status"` // READY, EXECUTED, ABORTED
	ExecutedAt      *time.Time `json:"executed_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
}
