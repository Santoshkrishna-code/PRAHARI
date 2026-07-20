package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in training lifecycle registers.
var transitionMatrix = map[Code][]Code{
	CodeDraft:      {CodePlanned},
	CodePlanned:    {CodeScheduled},
	CodeScheduled:  {CodeEnrollment},
	CodeEnrollment: {CodeInProgress},
	CodeInProgress: {CodeAssessment},
	CodeAssessment: {CodeCertified},
	CodeCertified:  {CodeActive},
	CodeActive:     {CodeRenewal},
	CodeRenewal:    {CodeActive},
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
