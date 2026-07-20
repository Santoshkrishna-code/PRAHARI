package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeDisconnected: {CodeConnecting},
	CodeConnecting:   {CodeConnected, CodeFailed, CodeDisconnected},
	CodeConnected:    {CodeDisconnected, CodeFailed},
	CodeFailed:       {CodeConnecting, CodeDisconnected},
}

// ValidateTransition checks if from → to state transition is allowed in Connector lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current Connector state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid Connector state transition from %s to %s", from, to)
}
