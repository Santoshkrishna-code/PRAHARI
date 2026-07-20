package restoration

import "time"

// Record tracks system restoration (removing locks/tags, de-isolating energy points).
type Record struct {
	ID            string    `json:"id"`
	CertificateID string    `json:"certificate_id"`
	RestoredBy    string    `json:"restored_by"`
	RestoredAt    time.Time `json:"restored_at"`
	Details       string    `json:"details"`
	ConfirmedSafe bool      `json:"confirmed_safe"`
}
