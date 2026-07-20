package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of a contractor onboarding.
var transitionMatrix = map[Code][]Code{
	CodeRegistered:    {CodeDocVerify},
	CodeDocVerify:     {CodeSafetyTrain},
	CodeSafetyTrain:   {CodeMedicalClear},
	CodeMedicalClear:  {CodeSiteInduction},
	CodeSiteInduction: {CodeApproved},
	CodeApproved:      {CodeActive},
	CodeActive:        {CodeSuspended, CodeExpired, CodeOffboarded},
	CodeSuspended:     {CodeActive, CodeOffboarded},
	CodeExpired:       {CodeActive, CodeOffboarded},
	CodeOffboarded:    {CodeArchived},
	CodeArchived:      {},
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
