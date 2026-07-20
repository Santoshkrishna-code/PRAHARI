package audit

import (
	"encoding/json"
	"time"
)

// Action defines user mutation tags.
type Action string

const (
	ActionCreated        Action = "CREATED"
	ActionUpdated        Action = "UPDATED"
	ActionDeleted        Action = "DELETED"
	ActionStatusChanged  Action = "STATUS_CHANGED"
	ActionApproved       Action = "APPROVED"
	ActionRejected       Action = "REJECTED"
	ActionRiskAssessed   Action = "RISK_ASSESSED"
	ActionGasTested      Action = "GAS_TESTED"
	ActionIsolated       Action = "ISOLATED"
	ActionCertIssued     Action = "CERT_ISSUED"
)

// Entry stores an immutable snapshot of mutations on domain tables.
type Entry struct {
	ID         string          `json:"id" db:"id"`
	EntityType string          `json:"entity_type" db:"entity_type"`
	EntityID   string          `json:"entity_id" db:"entity_id"`
	Action     Action          `json:"action" db:"action"`
	ActorID    string          `json:"actor_id" db:"actor_id"`
	OldValue   json.RawMessage `json:"old_value,omitempty" db:"old_value"`
	NewValue   json.RawMessage `json:"new_value,omitempty" db:"new_value"`
	IPAddress  string          `json:"ip_address,omitempty" db:"ip_address"`
	UserAgent  string          `json:"user_agent,omitempty" db:"user_agent"`
	Timestamp  time.Time       `json:"timestamp" db:"timestamp"`
}

// NewEntry constructs an entry record.
func NewEntry(entityType, entityID string, action Action, actorID string, oldVal, newVal interface{}) *Entry {
	entry := &Entry{
		EntityType: entityType,
		EntityID:   entityID,
		Action:     action,
		ActorID:    actorID,
		Timestamp:  time.Now(),
	}

	if oldVal != nil {
		entry.OldValue, _ = json.Marshal(oldVal)
	}
	if newVal != nil {
		entry.NewValue, _ = json.Marshal(newVal)
	}

	return entry
}
