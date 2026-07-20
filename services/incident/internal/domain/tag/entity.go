package tag

import (
	"errors"
)

// Tag represents a label that can be applied to incidents for classification and filtering.
type Tag struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Color string `json:"color" db:"color"`
}

// IncidentTag represents the many-to-many join between incidents and tags.
type IncidentTag struct {
	IncidentID string `json:"incident_id" db:"incident_id"`
	TagID      string `json:"tag_id" db:"tag_id"`
}

// Validate enforces domain invariants on the tag value object.
func (t *Tag) Validate() error {
	if t.Name == "" {
		return errors.New("tag name is required")
	}
	if len(t.Name) > 100 {
		return errors.New("tag name must not exceed 100 characters")
	}
	return nil
}
