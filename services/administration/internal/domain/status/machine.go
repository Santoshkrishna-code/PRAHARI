package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeDraft:       {CodeConfigured, CodeArchived},
	CodeConfigured:  {CodeValidated, CodeArchived},
	CodeValidated:   {CodeActivated, CodeArchived},
	CodeActivated:   {CodeOperational, CodeSuspended, CodeArchived},
	CodeOperational: {CodeSuspended, CodeArchived},
	CodeSuspended:   {CodeOperational, CodeArchived},
	CodeArchived:    {},
}

// ValidateTransition checks if from → to state transition is allowed in Organization lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current Organization state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid Organization state transition from %s to %s", from, to)
}
