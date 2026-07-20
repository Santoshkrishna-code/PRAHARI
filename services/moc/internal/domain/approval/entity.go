package approval

import "time"

// Record represents formal gate sign-off for MOC execution.
type Record struct {
	ID              string    `json:"id"`
	ChangeRequestID string    `json:"change_request_id"`
	ApproverID      string    `json:"approver_id"`
	Role            string    `json:"role"` // Plant Manager, PSM Lead, Engineering Lead
	Decision        string    `json:"decision"` // APPROVED, REJECTED
	Comments        string    `json:"comments"`
	ApprovedAt      time.Time `json:"approved_at"`
}
