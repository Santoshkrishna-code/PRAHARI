package status

import "fmt"

var transitionMatrix = map[Code][]Code{
	CodeDraft:    {CodeActive},
	CodeActive:   {CodeArchived},
	CodeArchived: {CodeActive},
}

// ValidateTransition checks if from → to state transition is allowed.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current twin state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid twin state transition from %s to %s", from, to)
}
