package documentapproval

import "time"

// Approval represents official sign-off for publication.
type Approval struct {
	ID         string    `json:"id"`
	DocumentID string    `json:"document_id"`
	VersionID  string    `json:"version_id"`
	ApproverID string    `json:"approver_id"`
	Approved   bool      `json:"approved"`
	ApprovedAt time.Time `json:"approved_at"`
	Comments   string    `json:"comments"`
}
