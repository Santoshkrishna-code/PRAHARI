package documentdistribution

import "time"

// Record tracks distribution list and acknowledgment of published documents.
type Record struct {
	ID             string     `json:"id"`
	DocumentID     string     `json:"document_id"`
	VersionID      string     `json:"version_id"`
	RecipientID    string     `json:"recipient_id"`
	DistributedAt  time.Time  `json:"distributed_at"`
	AcknowledgedAt *time.Time `json:"acknowledged_at,omitempty"`
}
