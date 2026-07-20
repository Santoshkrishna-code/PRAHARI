package retentionpolicy

import "time"

// Policy defines regulatory document retention and disposal rules.
type Policy struct {
	ID           string    `json:"id"`
	CategoryID   string    `json:"category_id"`
	DocumentType string    `json:"document_type"`
	RetentionYrs int       `json:"retention_years"`
	ActionOnExp  string    `json:"action_on_expiry"` // ARCHIVE, DESTROY, REVIEW
	CreatedAt    time.Time `json:"created_at"`
}
