package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodePlanned:               {CodeIsolationApproved, CodeCancelled},
	CodeIsolationApproved:     {CodeLocksApplied, CodeCancelled},
	CodeLocksApplied:          {CodeTagsApplied, CodeCancelled},
	CodeTagsApplied:           {CodeZeroEnergyVerified, CodeFailedVerification, CodeCancelled},
	CodeZeroEnergyVerified:    {CodeMaintenanceInProgress, CodeCancelled},
	CodeMaintenanceInProgress: {CodeRestorationPlanned},
	CodeRestorationPlanned:    {CodeLocksRemoved},
	CodeLocksRemoved:          {CodeReturnedToService},
	CodeReturnedToService:     {CodeClosed},
	CodeClosed:                {},
	CodeCancelled:             {},
	CodeFailedVerification:    {CodePlanned, CodeCancelled},
}

// ValidateTransition checks if from → to state transition is allowed in LOTO lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current LOTO state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid LOTO state transition from %s to %s", from, to)
}
