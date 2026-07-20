package template

import (
	"errors"
)

// Template represents dynamic localized message structures templates.
type Template struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Body string `json:"body"`
}

// Validate checks templates variables.
func (t *Template) Validate() error {
	if t.ID == "" {
		return errors.New("template ID is required")
	}
	if t.Name == "" {
		return errors.New("template Name is required")
	}
	return nil
}
