package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeRegistered:          {CodeAssigned, CodeRetired},
	CodeAssigned:            {CodeOperational, CodeRetired},
	CodeOperational:         {CodeInspection, CodeProofTest, CodeIntegrityAssessment, CodeBypassed, CodeImpaired, CodeOutOfService, CodeRetired},
	CodeInspection:          {CodeVerified, CodeImpaired, CodeOutOfService},
	CodeProofTest:           {CodeVerified, CodeImpaired, CodeOutOfService},
	CodeIntegrityAssessment: {CodeVerified, CodeImpaired, CodeOutOfService},
	CodeVerified:            {CodeOperational},
	CodeBypassed:            {CodeOperational, CodeImpaired},
	CodeImpaired:            {CodeOperational, CodeOutOfService},
	CodeOutOfService:        {CodeOperational, CodeRetired},
	CodeRetired:             {},
}

// ValidateTransition checks if from → to state transition is allowed in barrier lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current barrier state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid barrier state transition from %s to %s", from, to)
}

