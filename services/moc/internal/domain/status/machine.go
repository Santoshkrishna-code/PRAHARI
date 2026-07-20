package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeDraft:            {CodeImpactAssessment, CodeCancelled},
	CodeImpactAssessment: {CodeTechnicalReview, CodeRejected, CodeCancelled},
	CodeTechnicalReview:  {CodeRiskReview, CodeRejected, CodeCancelled},
	CodeRiskReview:       {CodeSafetyReview, CodeRejected, CodeCancelled},
	CodeSafetyReview:     {CodeApproval, CodeRejected, CodeCancelled},
	CodeApproval:         {CodeImplementation, CodeRejected, CodeCancelled},
	CodeImplementation:   {CodeVerification, CodeRolledBack, CodeCancelled},
	CodeVerification:     {CodeCloseout, CodeRolledBack, CodeCancelled},
	CodeCloseout:         {},
	CodeRejected:         {CodeDraft},
	CodeCancelled:        {},
	CodeRolledBack:       {CodeCloseout},
}

// ValidateTransition checks if from → to state transition is allowed in MOC lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current MOC state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid MOC transition from %s to %s", from, to)
}
