package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodePlanned:            {CodeScheduled, CodeCancelled},
	CodeScheduled:          {CodeInProgress, CodeRescheduled, CodeCancelled},
	CodeInProgress:         {CodeAttendanceRecorded, CodeCancelled},
	CodeAttendanceRecorded: {CodeMinutesApproved, CodeCancelled},
	CodeMinutesApproved:    {CodeActionsGenerated, CodeClosed},
	CodeActionsGenerated:   {CodeClosed},
	CodeClosed:             {},
	CodeCancelled:          {},
	CodeRescheduled:        {CodeScheduled, CodeCancelled},
}

// ValidateTransition checks if from → to state transition is allowed in Meeting lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current Meeting state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid Meeting state transition from %s to %s", from, to)
}
