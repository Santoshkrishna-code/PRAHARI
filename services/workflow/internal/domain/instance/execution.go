package instance

import (
	"errors"
)

// Instance tracks execution variables, state values, and running step IDs.
type Instance struct {
	ID            string            `json:"id"`
	DefinitionID  string            `json:"definition_id"`
	Version       int               `json:"version"`
	CurrentStepID string            `json:"current_step_id"`
	State         string            `json:"state"` // e.g. "RUNNING", "COMPLETED"
	Variables     map[string]string `json:"variables"`
}

// Validate checks entity values.
func (i *Instance) Validate() error {
	if i.ID == "" {
		return errors.New("instance ID is required")
	}
	if i.DefinitionID == "" {
		return errors.New("instance associated definition ID is required")
	}
	return nil
}
