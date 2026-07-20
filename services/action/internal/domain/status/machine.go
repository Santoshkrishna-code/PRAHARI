package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeCreated:             {CodeAssigned, CodeCancelled},
	CodeAssigned:            {CodeInProgress, CodeOverdue, CodeCancelled},
	CodeInProgress:          {CodeEvidenceSubmitted, CodeOverdue, CodeCancelled},
	CodeEvidenceSubmitted:   {CodeEffectivenessReview, CodeRejected},
	CodeEffectivenessReview: {CodeClosed, CodeRejected},
	CodeClosed:              {},
	CodeCancelled:           {},
	CodeOverdue:             {CodeInProgress, CodeCancelled},
	CodeRejected:            {CodeInProgress, CodeCancelled},
}

// ValidateTransition checks if from → to state transition is allowed in Action lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current Action state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid Action state transition from %s to %s", from, to)
}
