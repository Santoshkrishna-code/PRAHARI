package replacement

import "time"

// Plan outlines automated scheduling of PPE replacement based on expiry dates or wear & tear failures.
type Plan struct {
	ID          string    `json:"id"`
	ItemID      string    `json:"item_id"`
	TriggerReason string  `json:"trigger_reason"` // EXPIRED, INSPECTION_FAILED, SCHEDULED_REPLACEMENT
	ScheduledAt time.Time `json:"scheduled_at"`
	ReplacedBy  string    `json:"replaced_by,omitempty"`
	ReplacedAt  *time.Time `json:"replaced_at,omitempty"`
	Status      string    `json:"status"` // PENDING, REPLACED, CANCELLED
}
