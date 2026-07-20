package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of an inspection.
var transitionMatrix = map[Code][]Code{
	CodeDraft:       {CodeScheduled, CodeCancelled},
	CodeScheduled:   {CodeAssigned, CodeCancelled},
	CodeAssigned:    {CodeInProgress},
	CodeInProgress:  {CodeCompleted},
	CodeCompleted:   {CodeUnderReview},
	CodeUnderReview: {CodeApproved, CodeRejected},
	CodeApproved:    {CodeClosed},
	CodeRejected:    {CodeAssigned}, // Rejected inspections get reassigned/re-executed
	CodeClosed:      {CodeArchived},
	CodeArchived:    {},
	CodeCancelled:   {},
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
