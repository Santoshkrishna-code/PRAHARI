package saga

import (
	"errors"
)

// CompensableAction represents reverse rollback instructions.
type CompensableAction struct {
	StepID           string `json:"step_id"`
	CompensateStepID string `json:"compensate_step_id"`
}

// Validate checks fields.
func (c *CompensableAction) Validate() error {
	if c.StepID == "" || c.CompensateStepID == "" {
		return errors.New("Saga actions require execution and compensation steps IDs mapping")
	}
	return nil
}
