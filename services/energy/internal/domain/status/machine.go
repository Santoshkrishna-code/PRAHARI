package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeMeterRegistered: {CodeDataCollection, CodeArchived},
	CodeDataCollection:  {CodeValidation, CodeArchived},
	CodeValidation:      {CodeAggregation, CodeArchived},
	CodeAggregation:     {CodeAnalysis, CodeArchived},
	CodeAnalysis:        {CodeOptimization, CodeReporting, CodeArchived},
	CodeOptimization:    {CodeReporting, CodeArchived},
	CodeReporting:       {CodeArchived},
	CodeArchived:        {CodeMeterRegistered},
}

// ValidateTransition checks if from → to state transition is allowed.
func ValidateTransition(from, to Code) error {
	allowed, exists := transitionMatrix[from]
	if !exists {
		return fmt.Errorf("unknown Energy state: %s", from)
	}

	for _, t := range allowed {
		if t == to {
			return nil
		}
	}

	return fmt.Errorf("invalid Energy state transition: %s → %s", from, to)
}
