package documentverification

import "time"

// Record tracks verification of visitor credentials, certs, and agreements (e.g. NDA, Liability waiver).
type Record struct {
	ID             string    `json:"id"`
	VisitorID      string    `json:"visitor_id"`
	DocumentName   string    `json:"document_name"` // NDA, Liability waiver, Govt ID
	Verified       bool      `json:"verified"`
	VerifiedBy     string    `json:"verified_by"`
	VerifiedAt     time.Time `json:"verified_at"`
	DocumentDocRef string    `json:"document_doc_ref"` // reference to Document Management
}
