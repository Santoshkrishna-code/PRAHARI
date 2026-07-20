package task

import (
	"errors"
)

// Task represents an allocated operational workflow assignment.
type Task struct {
	ID         string `json:"id"`
	InstanceID string `json:"instance_id"`
	Name       string `json:"name"`
	Status     string `json:"status"` // e.g. "pending", "completed"
	AssignedTo string `json:"assigned_to"`
}

// Validate checks task structures.
func (t *Task) Validate() error {
	if t.ID == "" {
		return errors.New("task ID is required")
	}
	if t.InstanceID == "" {
		return errors.New("associated instance ID is required")
	}
	return nil
}
