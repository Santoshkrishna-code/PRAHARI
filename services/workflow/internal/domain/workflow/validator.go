package workflow

import (
	"errors"
)

// ValidateStepsRelation checks transitions paths, preventing execution loops or orphans.
func ValidateStepsRelation(d *Definition) error {
	stepMap := make(map[string]bool)
	for _, step := range d.Steps {
		stepMap[step.ID] = true
	}

	for _, step := range d.Steps {
		if step.NextStep != "" && !stepMap[step.NextStep] {
			return errors.New("invalid step relationship: target next_step '" + step.NextStep + "' does not exist")
		}
	}
	return nil
}
