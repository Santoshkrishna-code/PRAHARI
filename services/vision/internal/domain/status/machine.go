package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeOnline:      {CodeOffline, CodeUnreachable},
	CodeOffline:     {CodeOnline, CodeUnreachable},
	CodeUnreachable: {CodeOnline, CodeOffline},
}

// ValidateTransition checks if from → to state transition is allowed in Camera connection status.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current Camera state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid Camera state transition from %s to %s", from, to)
}
