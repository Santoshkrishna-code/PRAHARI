package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeQueued:     {CodeParsing, CodeFailed},
	CodeParsing:    {CodeVectorized, CodeFailed},
	CodeVectorized: {},
	CodeFailed:     {CodeQueued},
}

// ValidateTransition checks if from → to state transition is allowed in document indexing pipeline.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current document state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid document state transition from %s to %s", from, to)
}
