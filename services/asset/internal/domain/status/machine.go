package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of an asset.
var transitionMatrix = map[Code][]Code{
	CodeRegistered:     {CodeCommissioned},
	CodeCommissioned:   {CodeOperational, CodeOutOfService},
	CodeOperational:    {CodeMaintenance, CodeOutOfService},
	CodeMaintenance:    {CodeOperational, CodeOutOfService},
	CodeOutOfService:   {CodeOperational, CodeMaintenance, CodeDecommissioned},
	CodeDecommissioned: {CodeDisposed},
	CodeDisposed:       {},
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

	return fmt.Errorf("invalid lifecycle state transition: %s → %s is not permitted", from, to)
}
