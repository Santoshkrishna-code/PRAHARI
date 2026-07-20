package status

import (
	"fmt"
)

// transitionMatrix defines allowed status flows in the lifecycle of a near miss report.
var transitionMatrix = map[Code][]Code{
	CodeReported:    {CodeClassified, CodeEscalated},
	CodeClassified:  {CodeInvestigate, CodeEscalated},
	CodeInvestigate: {CodeRootCause, CodeEscalated},
	CodeRootCause:   {CodeCorrective, CodeEscalated},
	CodeCorrective:  {CodeVerified},
	CodeVerified:    {CodeClosed},
	CodeClosed:      {},
	CodeEscalated:   {},
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
