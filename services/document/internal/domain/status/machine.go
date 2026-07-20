package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeDraft:                  {CodeReview, CodeWithdrawn},
	CodeReview:                 {CodeApproval, CodeRejected, CodeDraft},
	CodeApproval:               {CodePublished, CodeRejected, CodeDraft},
	CodePublished:              {CodeControlledDistribution, CodePeriodicReview, CodeRevision, CodeSuperseded, CodeArchived, CodeWithdrawn},
	CodeControlledDistribution: {CodePeriodicReview, CodeRevision, CodeSuperseded, CodeArchived},
	CodePeriodicReview:         {CodeRevision, CodePublished, CodeSuperseded, CodeArchived},
	CodeRevision:               {CodeDraft, CodeReview},
	CodeSuperseded:             {CodeArchived},
	CodeArchived:               {},
	CodeRejected:               {CodeDraft, CodeWithdrawn},
	CodeWithdrawn:              {},
}

// ValidateTransition checks if from → to state transition is allowed in Document lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current document state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid document state transition from %s to %s", from, to)
}
