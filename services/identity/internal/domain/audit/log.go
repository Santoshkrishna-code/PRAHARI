package audit

import (
	"time"
)

// AuditLog tracks core system mutations and action timestamps.
type AuditLog struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Action    string    `json:"action"` // e.g. "identity.user.created"
	ClientIP  string    `json:"client_ip"`
	Timestamp time.Time `json:"timestamp"`
}
