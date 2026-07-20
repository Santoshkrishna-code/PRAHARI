package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeSourceRegistered:  {CodeCollection},
	CodeCollection:        {CodeStorage, CodeTreatment},
	CodeStorage:           {CodeTreatment, CodeDistribution},
	CodeTreatment:         {CodeDistribution, CodeStorage},
	CodeDistribution:      {CodeConsumption},
	CodeConsumption:       {CodeRecyclingReuse, CodePerformanceReview},
	CodeRecyclingReuse:    {CodePerformanceReview, CodeStorage, CodeDistribution},
	CodePerformanceReview: {CodeArchived, CodeSourceRegistered},
	CodeArchived:          {CodeSourceRegistered},
}

// ValidateTransition checks if a state change is allowed according to the water lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid transition from %s to %s", from, to)
}
