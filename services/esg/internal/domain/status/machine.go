package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeObjectiveDefined:    {CodeBaselineEstablished, CodeArchived},
	CodeBaselineEstablished: {CodeDataCollection, CodeArchived},
	CodeDataCollection:      {CodeCalculation, CodeArchived},
	CodeCalculation:         {CodeValidation, CodeArchived},
	CodeValidation:          {CodeDisclosure, CodeArchived},
	CodeDisclosure:          {CodeExecutiveReview, CodeArchived},
	CodeExecutiveReview:     {CodePublished, CodeReopened, CodeArchived},
	CodePublished:           {CodeReopened, CodeArchived},
	CodeReopened:            {CodeDataCollection, CodeCalculation},
	CodeArchived:            {CodeObjectiveDefined},
}

// ValidateTransition checks if from → to state transition is allowed.
func ValidateTransition(from, to Code) error {
	allowed, exists := transitionMatrix[from]
	if !exists {
		return fmt.Errorf("unknown ESG state: %s", from)
	}

	for _, t := range allowed {
		if t == to {
			return nil
		}
	}

	return fmt.Errorf("invalid ESG state transition: %s → %s", from, to)
}
