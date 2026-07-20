package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeAvailable:   {CodeReserved, CodeIssued, CodeExpired, CodeDisposed},
	CodeReserved:    {CodeIssued, CodeAvailable, CodeExpired},
	CodeIssued:      {CodeInUse, CodeReturned, CodeExpired},
	CodeInUse:       {CodeInspection, CodeReturned, CodeExpired},
	CodeInspection:  {CodeReturned, CodeMaintenance, CodeDisposed},
	CodeReturned:    {CodeAvailable, CodeInspection, CodeDisposed},
	CodeMaintenance: {CodeAvailable, CodeDisposed},
	CodeExpired:     {CodeDisposed},
	CodeDisposed:    {},
}

// ValidateTransition checks if from → to state transition is allowed in PPE lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current ppe item state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid ppe item state transition from %s to %s", from, to)
}
