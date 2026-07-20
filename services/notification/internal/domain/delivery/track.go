package delivery

// State definitions tracking active lifecycle states.
const (
	StateQueued    = "QUEUED"
	StateSent      = "SENT"
	StateRead      = "READ"
	StateFailed    = "FAILED"
	StateCancelled = "CANCELLED"
)

// Track models outbound message lifecycle histories.
type Track struct {
	NotificationID string `json:"notification_id"`
	State          string `json:"state"`
}
