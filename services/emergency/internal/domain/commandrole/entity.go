package commandrole

import "time"

// Role represents an ICS position assignment (Incident Commander, Safety Officer, Operations Chief, Logistics Chief, Liaison Officer).
type Role struct {
	ID          string    `json:"id"`
	CommandID   string    `json:"command_id"`
	RoleName    string    `json:"role_name"`
	AssigneeID  string    `json:"assignee_id"`
	AssignedAt  time.Time `json:"assigned_at"`
}
