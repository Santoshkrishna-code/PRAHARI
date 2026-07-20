package documentrelationship

import "time"

// Relation maps linkages between controlled documents (e.g., SOP references P&ID).
type Relation struct {
	ID           string    `json:"id"`
	SourceDocID  string    `json:"source_doc_id"`
	TargetDocID  string    `json:"target_doc_id"`
	RelationType string    `json:"relation_type"` // REFERENCES, SUPERSEDES, DERIVED_FROM, ATTACHMENT_TO
	CreatedAt    time.Time `json:"created_at"`
}
