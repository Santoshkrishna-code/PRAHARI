package documentupdate

import "time"

// Record tracks P&ID, SOP, and engineering documentation revisions triggered by MOC.
type Record struct {
	ID              string    `json:"id"`
	ChangeRequestID string    `json:"change_request_id"`
	DocumentNumber  string    `json:"document_number"`
	DocumentTitle   string    `json:"document_title"`
	NewRevision     string    `json:"new_revision"`
	IsUpdated       bool      `json:"is_updated"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at"`
}
