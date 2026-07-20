package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeDraft:                  {CodePreparation, CodeCancelled},
	CodePreparation:            {CodeStudy, CodeCancelled},
	CodeStudy:                  {CodeRiskEvaluation, CodeCancelled},
	CodeRiskEvaluation:         {CodeRecommendation, CodeCancelled},
	CodeRecommendation:         {CodeApproval, CodeCancelled},
	CodeApproval:               {CodeImplementationTracking, CodeCancelled},
	CodeImplementationTracking: {CodeVerification, CodeCancelled},
	CodeVerification:           {CodeRevalidationScheduled, CodeClosed, CodeCancelled},
	CodeRevalidationScheduled:  {CodeClosed, CodeSuperseded},
	CodeClosed:                 {CodeSuperseded},
	CodeCancelled:              {},
	CodeSuperseded:             {},
}

// ValidateTransition checks if from → to state transition is allowed in PHA study lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current PHA state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid PHA transition from %s to %s", from, to)
}
