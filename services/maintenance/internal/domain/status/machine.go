package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of a maintenance work order.
var transitionMatrix = map[Code][]Code{
	CodeDraft:       {CodePlanned, CodeCancelled},
	CodePlanned:     {CodeApproved, CodeCancelled},
	CodeApproved:    {CodeScheduled, CodeCancelled},
	CodeScheduled:   {CodeAssigned, CodeCancelled},
	CodeAssigned:    {CodeInProgress},
	CodeInProgress:  {CodeCompleted},
	CodeCompleted:   {CodeVerified},
	CodeVerified:    {CodeClosed},
	CodeClosed:      {},
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
