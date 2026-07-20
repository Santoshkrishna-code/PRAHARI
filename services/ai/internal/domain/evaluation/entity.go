package evaluation

import "time"

// Score tracks calculated generation relevance quality.
type Score struct {
	ID         string    `json:"id"`
	ResponseID string    `json:"response_id"`
	Relevance  float64   `json:"relevance"` // E.g., 0.0 to 1.0
	Accuracy   float64   `json:"accuracy"`
	TestedAt   time.Time `json:"tested_at"`
}
