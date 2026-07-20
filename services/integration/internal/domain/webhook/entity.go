package webhook

import "time"

// Subscription represents an outbound HTTP webhook subscription registered by client systems.
type Subscription struct {
	ID        string    `json:"id"`
	TargetURL string    `json:"target_url"`
	EventName string    `json:"event_name"` // E.g., incident.created
	SecretKey string    `json:"secret_key,omitempty"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}
