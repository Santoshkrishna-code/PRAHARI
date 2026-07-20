package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of a permit.
var transitionMatrix = map[Code][]Code{
	CodeDraft:          {CodeSubmitted, CodeCancelled},
	CodeSubmitted:      {CodeRiskAssessment, CodeRejected, CodeCancelled},
	CodeRiskAssessment: {CodeApproval, CodeRejected},
	CodeApproval:       {CodeIssued, CodeRejected},
	CodeIssued:         {CodeActive},
	CodeActive:         {CodeSuspended, CodeCompleted},
	CodeSuspended:      {CodeActive, CodeCancelled},
	CodeCompleted:      {CodeClosed},
	CodeClosed:         {CodeArchived},
	CodeArchived:       {},
	CodeRejected:       {},
	CodeCancelled:      {},
}

// ValidateTransition checks if a status transition is permitted.
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

// GetAllowedTransitions returns transitions from a source state.
func GetAllowedTransitions(from Code) []Code {
	return transitionMatrix[from]
}
