package checklistitem

import (
	"errors"
)

// ResponseType defines response constraint categories.
type ResponseType string

const (
	ResponsePassFail    ResponseType = "PASS_FAIL"
	ResponseYesNo       ResponseType = "YES_NO"
	ResponseMeasurement ResponseType = "MEASUREMENT"
)

// ChecklistItem is the individual question check item response record.
type ChecklistItem struct {
	ID            string       `json:"id" db:"id"`
	ChecklistID   string       `json:"checklist_id" db:"checklist_id"`
	Question      string       `json:"question" db:"question"`
	Description   string       `json:"description" db:"description"`
	CategoryName  string       `json:"category_name" db:"category_name"`
	ResponseType  ResponseType `json:"response_type" db:"response_type"`
	ResponseValue string       `json:"response_value" db:"response_value"`
	IsPassed      bool         `json:"is_passed" db:"is_passed"`
	Comments      string       `json:"comments,omitempty" db:"comments"`
}

// Validate checks domain invariants.
func (ci *ChecklistItem) Validate() error {
	if ci.ChecklistID == "" {
		return errors.New("checklist ID is required for checklist item")
	}
	if ci.Question == "" {
		return errors.New("checklist item question is required")
	}
	return nil
}
