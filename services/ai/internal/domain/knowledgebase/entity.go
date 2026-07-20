package knowledgebase

import "time"

// Source represents external sources synced into vector directories.
type Source struct {
	ID         string    `json:"id"`
	SourceName string    `json:"source_name"` // E.g., LOTO procedures, Chemical SDS
	Status     string    `json:"status"`      // SYNCED, OUTDATED
	UpdatedAt  time.Time `json:"updated_at"`
}
