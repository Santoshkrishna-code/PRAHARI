package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of a risk assessment register record.
var transitionMatrix = map[Code][]Code{
	CodeDraft:      {CodeAssessment},
	CodeAssessment: {CodeReview},
	CodeReview:     {CodeApproval},
	CodeApproval:   {CodeActive},
	CodeActive:     {CodeReassess, CodeClosed},
	CodeReassess:   {CodeReview, CodeActive},
	CodeClosed:     {},
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
