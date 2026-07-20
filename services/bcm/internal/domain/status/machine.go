package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodePlanned:                {CodeBusinessImpactAnalysis, CodeArchived},
	CodeBusinessImpactAnalysis: {CodeStrategyDevelopment, CodeArchived},
	CodeStrategyDevelopment:    {CodePlanDevelopment, CodeArchived},
	CodePlanDevelopment:        {CodeApproval, CodeArchived},
	CodeApproval:               {CodeExercise, CodeActivation, CodeArchived},
	CodeExercise:               {CodeApproval, CodeActivation, CodeArchived},
	CodeActivation:             {CodeRecovery, CodeArchived},
	CodeRecovery:               {CodeReview, CodeArchived},
	CodeReview:                 {CodeContinuousImprovement, CodeArchived},
	CodeContinuousImprovement:  {CodeApproval, CodeArchived},
	CodeArchived:               {},
}

// ValidateTransition checks if from → to state transition is allowed in BCM lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current BCM state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid BCM state transition from %s to %s", from, to)
}
