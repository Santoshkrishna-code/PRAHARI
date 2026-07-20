package checklisttemplate

import (
	"encoding/json"
	"errors"
)

// ChecklistTemplate defines a reusable question checklist template format.
type ChecklistTemplate struct {
	ID          string          `json:"id" db:"id"`
	Name        string          `json:"name" db:"name"`
	Description string          `json:"description" db:"description"`
	Categories  json.RawMessage `json:"categories" db:"categories"` // JSON array of category names
	Items       json.RawMessage `json:"items" db:"items"`           // JSON array of predefined template items
	IsActive    bool            `json:"is_active" db:"is_active"`
}

// Validate checks domain invariants for ChecklistTemplate.
func (ct *ChecklistTemplate) Validate() error {
	if ct.Name == "" {
		return errors.New("checklist template name is required")
	}
	return nil
}
