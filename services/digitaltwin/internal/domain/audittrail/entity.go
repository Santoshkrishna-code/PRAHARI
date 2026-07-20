package audittrail

import "time"

// Record tracks digital twin configuration and state changes.
type Record struct {
	ID         string    `json:"id"`
	Action     string    `json:"action"`
	Resource   string    `json:"resource"`
	ResourceID string    `json:"resource_id"`
	ActorID    string    `json:"actor_id"`
	Timestamp  time.Time `json:"timestamp"`
	OldState   string    `json:"old_state,omitempty"`
	NewState   string    `json:"new_state,omitempty"`
}
