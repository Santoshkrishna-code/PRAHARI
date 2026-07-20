package checklist

import (
	"errors"
)

// Checklist holds actual execution response checklist templates linked to an inspection.
type Checklist struct {
	ID                  string `json:"id" db:"id"`
	InspectionID        string `json:"inspection_id" db:"inspection_id"`
	ChecklistTemplateID string `json:"checklist_template_id" db:"checklist_template_id"`
	Name                string `json:"name" db:"name"`
}

// Validate checks domain invariants.
func (c *Checklist) Validate() error {
	if c.InspectionID == "" {
		return errors.New("inspection ID is required for checklist")
	}
	if c.Name == "" {
		return errors.New("checklist name is required")
	}
	return nil
}
