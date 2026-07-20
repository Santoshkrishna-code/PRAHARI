package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodePrepared:           {CodeDetected, CodeDeclared, CodeCancelled},
	CodeDetected:           {CodeDeclared, CodeCancelled},
	CodeDeclared:           {CodeResponseActivated, CodeCancelled},
	CodeResponseActivated:  {CodeCommandEstablished, CodeCancelled},
	CodeCommandEstablished: {CodeResourceDeployment, CodeEvacuation, CodeStabilized, CodeCancelled},
	CodeResourceDeployment: {CodeEvacuation, CodeStabilized, CodeCancelled},
	CodeEvacuation:         {CodeStabilized, CodeCancelled},
	CodeStabilized:         {CodeRecovery, CodeAfterActionReview, CodeClosed},
	CodeRecovery:           {CodeAfterActionReview, CodeClosed},
	CodeAfterActionReview:  {CodeClosed},
	CodeClosed:             {},
	CodeCancelled:          {},
}

// ValidateTransition checks if from → to state transition is allowed in emergency lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current emergency state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid emergency state transition from %s to %s", from, to)
}
