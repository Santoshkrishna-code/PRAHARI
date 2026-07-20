package minutes

import "time"

// Minutes represents the approved record of discussions and decisions from a meeting.
type Minutes struct {
	ID          string     `json:"id"`
	MeetingID   string     `json:"meeting_id"`
	Body        string     `json:"body"`
	RecorderID  string     `json:"recorder_id"`
	ApproverID  string     `json:"approver_id,omitempty"`
	ApprovedAt  *time.Time `json:"approved_at,omitempty"`
	Status      string     `json:"status"` // DRAFT, SUBMITTED, APPROVED
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
