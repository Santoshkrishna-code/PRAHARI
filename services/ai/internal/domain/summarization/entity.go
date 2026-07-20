package summarization

import "time"

// Summary represents generated condensed reports of incident history.
type Summary struct {
	ID        string    `json:"id"`
	SourceID  string    `json:"source_id"`
	Original  string    `json:"original"`
	Condensed string    `json:"condensed"`
	CreatedAt time.Time `json:"created_at"`
}
