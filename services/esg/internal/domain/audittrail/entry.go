package audittrail

import (
	"time"
)

// Entry maps immutable logs.
type Entry struct {
	ID         string            `json:"id" db:"id"`
	Action     string            `json:"action" db:"action"`
	Resource   string            `json:"resource" db:"resource"`
	ResourceID string            `json:"resource_id" db:"resource_id"`
	ActorID    string            `json:"actor_id" db:"actor_id"`
	Timestamp  time.Time         `json:"timestamp" db:"timestamp"`
	OldState   map[string]string `json:"old_state" db:"old_state"`
	NewState   map[string]string `json:"new_state" db:"new_state"`
}

// NewEntry constructs Entry.
func NewEntry(action, res, resID, actor string, oldS, newS map[string]string) *Entry {
	return &Entry{
		Action:     action,
		Resource:   res,
		ResourceID: resID,
		ActorID:    actor,
		Timestamp:  time.Now(),
		OldState:   oldS,
		NewState:   newS,
	}
}
