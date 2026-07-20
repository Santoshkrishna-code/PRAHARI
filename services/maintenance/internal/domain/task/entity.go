package task

import (
	"errors"
)

// Task represents an individual operation step within a work order execution checklist.
type Task struct {
	ID            string `json:"id" db:"id"`
	MaintenanceID string `json:"maintenance_id" db:"maintenance_id"`
	Description   string `json:"description" db:"description"`
	SequenceOrder int    `json:"sequence_order" db:"sequence_order"`
	IsCompleted   bool   `json:"is_completed" db:"is_completed"`
}

// Validate checks domain invariants.
func (t *Task) Validate() error {
	if t.MaintenanceID == "" {
		return errors.New("maintenance ID is required")
	}
	if t.Description == "" {
		return errors.New("task description cannot be empty")
	}
	return nil
}
