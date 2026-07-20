package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in compliance obligations lifecycle registers.
var transitionMatrix = map[Code][]Code{
	CodeDraft:      {CodeDefined},
	CodeDefined:    {CodeAssigned},
	CodeAssigned:   {CodeEvidence},
	CodeEvidence:   {CodeReview},
	CodeReview:     {CodeCompliant},
	CodeCompliant:  {CodeMonitoring},
	CodeMonitoring: {CodeRenewal},
	CodeRenewal:    {CodeCompliant},
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
