package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of a hazard report.
var transitionMatrix = map[Code][]Code{
	CodeIdentified:  {CodeReported, CodeRejected},
	CodeReported:    {CodeAssessed, CodeRejected},
	CodeAssessed:    {CodeMitigating, CodeRejected},
	CodeMitigating:  {CodeApproved, CodeRejected},
	CodeApproved:    {CodeImplemented},
	CodeImplemented: {CodeVerified},
	CodeVerified:    {CodeClosed},
	CodeClosed:      {},
	CodeRejected:    {},
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
