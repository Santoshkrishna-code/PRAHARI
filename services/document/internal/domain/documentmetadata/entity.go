package documentmetadata

import "time"

// Attribute stores key-value pairs for flexible metadata indexing & enterprise search.
type Attribute struct {
	ID         string    `json:"id"`
	DocumentID string    `json:"document_id"`
	MetaKey    string    `json:"meta_key"`
	MetaValue  string    `json:"meta_value"`
	CreatedAt  time.Time `json:"created_at"`
}
