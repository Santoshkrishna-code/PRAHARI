package approval

import (
	"errors"
)

// Approval maps multi-approver structures.
type Approval struct {
	TaskID    string            `json:"task_id"`
	Approvers []string          `json:"approvers"`
	Strategy  string            `json:"strategy"` // "sequential", "quorum", "any"
	Decisions map[string]string `json:"decisions"` // approver_id -> "approve"/"reject"
}

// Validate checks entity values.
func (a *Approval) Validate() error {
	if a.TaskID == "" {
		return errors.New("approval associated task ID is required")
	}
	if len(a.Approvers) == 0 {
		return errors.New("approvers list cannot be empty")
	}
	return nil
}
