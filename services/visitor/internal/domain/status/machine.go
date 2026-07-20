package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeScheduled:            {CodeHostApproval, CodeCancelled, CodeBlacklisted},
	CodeHostApproval:         {CodeSecurityVerification, CodeRejected, CodeCancelled, CodeBlacklisted},
	CodeSecurityVerification: {CodeGatePassIssued, CodeRejected, CodeCancelled, CodeBlacklisted},
	CodeGatePassIssued:       {CodeCheckedIn, CodeCancelled, CodeBlacklisted},
	CodeCheckedIn:            {CodeOnSite, CodeCheckedOut, CodeBlacklisted},
	CodeOnSite:               {CodeCheckedOut, CodeBlacklisted},
	CodeCheckedOut:           {CodeClosed},
	CodeClosed:               {},
	CodeRejected:             {},
	CodeCancelled:            {},
	CodeBlacklisted:          {},
}

// ValidateTransition checks if from → to state transition is allowed in Visitor lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current visitor state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid visitor state transition from %s to %s", from, to)
}
