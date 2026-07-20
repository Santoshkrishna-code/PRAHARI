package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeDraft:     {CodeGenerated, CodeFailed},
	CodeGenerated: {CodeDelivered, CodeFailed},
	CodeDelivered: {},
	CodeFailed:    {CodeDraft},
}

// ValidateTransition checks if from → to state transition is allowed in Report lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current Report state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid Report state transition from %s to %s", from, to)
}
