package status

import (
	"fmt"
)

// transitionMatrix defines the set of valid state transitions in the incident lifecycle.
// Each key maps a source state to the set of target states it may transition to.
var transitionMatrix = map[Code][]Code{
	CodeDraft:          {CodeSubmitted, CodeCancelled},
	CodeSubmitted:      {CodeUnderReview, CodeRejected, CodeCancelled},
	CodeUnderReview:    {CodeAssigned, CodeRejected},
	CodeAssigned:       {CodeInvestigating},
	CodeInvestigating:  {CodeCAPAInProgress, CodeResolved},
	CodeCAPAInProgress: {CodeResolved},
	CodeResolved:       {CodeClosed},
	// Terminal states have no outgoing transitions.
	CodeClosed:    {},
	CodeRejected:  {},
	CodeCancelled: {},
}

// ValidateTransition enforces the state machine rules. Returns an error if the
// transition from 'from' to 'to' is not allowed by the lifecycle definition.
func ValidateTransition(from, to Code) error {
	allowedTargets, exists := transitionMatrix[from]
	if !exists {
		return fmt.Errorf("unknown source state: %s", from)
	}

	for _, allowed := range allowedTargets {
		if to == allowed {
			return nil
		}
	}

	return fmt.Errorf("invalid state transition: %s → %s is not permitted", from, to)
}

// GetAllowedTransitions returns the set of valid target states for a given source state.
func GetAllowedTransitions(from Code) []Code {
	targets, exists := transitionMatrix[from]
	if !exists {
		return nil
	}
	return targets
}
