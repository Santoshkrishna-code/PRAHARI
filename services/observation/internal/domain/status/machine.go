package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of a safety observation report.
var transitionMatrix = map[Code][]Code{
	CodeObserved: {CodeRecorded},
	CodeRecorded: {CodeReviewed},
	CodeReviewed: {CodeCoaching, CodeClosed},
	CodeCoaching: {CodeFollowUp},
	CodeFollowUp: {CodeVerified},
	CodeVerified: {CodeClosed},
	CodeClosed:   {},
}

// ValidateTransition checks transition paths constraints.
func ValidateTransition(from, to Code) error {
	allowed, exists := transitionMatrix[from]
	if !exists {
		return fmt.Errorf("unknown source state: %s", from)
	}

	for _, t := range allowed {
		if t == to {
			return nil
		}
	}

	return fmt.Errorf("invalid state transition: %s → %s is not permitted", from, to)
}
