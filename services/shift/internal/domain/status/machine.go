package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeScheduled:         {CodeCrewAssigned, CodeCancelled},
	CodeCrewAssigned:      {CodeShiftStarted, CodeCancelled},
	CodeShiftStarted:      {CodeOperational, CodeCancelled},
	CodeOperational:       {CodeHandoverInitiated, CodeShiftClosed},
	CodeHandoverInitiated: {CodeHandoverAccepted, CodeOperational},
	CodeHandoverAccepted:  {CodeShiftClosed},
	CodeShiftClosed:       {},
	CodeCancelled:         {},
}

// ValidateTransition checks if from → to state transition is allowed in Shift lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current shift state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid shift state transition from %s to %s", from, to)
}
