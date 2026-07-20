package prompt

import "time"

// Template holds versioned system instructions for prompts.
type Template struct {
	ID        string    `json:"id"`
	Key       string    `json:"key"` // E.g., rca_analysis, sds_lookup
	Version   int       `json:"version"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}
