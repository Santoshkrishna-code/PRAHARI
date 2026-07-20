package checklistitem

import (
	"errors"
)

// ChecklistItem maps specific compliance questions.
type ChecklistItem struct {
	ID          string `json:"id" db:"id"`
	ChecklistID string `json:"checklist_id" db:"checklist_id"`
	Question    string `json:"question" db:"question"`
}

// Validate checks domain invariants.
func (ci *ChecklistItem) Validate() error {
	if ci.ChecklistID == "" {
		return errors.New("checklist ID reference is required")
	}
	if ci.Question == "" {
		return errors.New("question text cannot be empty")
	}
	return nil
}
