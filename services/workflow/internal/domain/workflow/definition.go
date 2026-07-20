package workflow

import (
	"errors"
)

// Step represents a single executable action inside the DSL workflow structure.
type Step struct {
	ID       string `json:"id" yaml:"id"`
	Type     string `json:"type" yaml:"type"` // e.g. "approval", "webhook", "email"
	NextStep string `json:"next_step" yaml:"next_step"`
}

// Definition defines the versioned, immutable template blueprint schema.
type Definition struct {
	ID      string `json:"id" yaml:"id"`
	Name    string `json:"name" yaml:"name"`
	Version int    `json:"version" yaml:"version"`
	Steps   []Step `json:"steps" yaml:"steps"`
}

// Validate asserts completeness of workflow schemas.
func (d *Definition) Validate() error {
	if d.ID == "" {
		return errors.New("workflow definition ID is required")
	}
	if d.Name == "" {
		return errors.New("workflow definition Name is required")
	}
	if len(d.Steps) == 0 {
		return errors.New("workflow definition must contain at least one step block")
	}
	return nil
}
