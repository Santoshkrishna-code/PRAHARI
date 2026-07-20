package feedback

import "time"

// Item logs users thumbs-up or down votes for quality loop inputs.
type Item struct {
	ID         string    `json:"id"`
	ResponseID string    `json:"response_id"`
	Vote       int       `json:"vote"` // 1 (Up), -1 (Down)
	Comment    string    `json:"comment,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
