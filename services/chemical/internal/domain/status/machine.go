package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeRequested:           {CodeTechnicalReview, CodeRejected},
	CodeTechnicalReview:     {CodeSafetyReview, CodeRejected},
	CodeSafetyReview:        {CodeEnvironmentalReview, CodeRejected},
	CodeEnvironmentalReview: {CodeApproved, CodeRejected},
	CodeApproved:            {CodePurchased, CodeRecalled},
	CodePurchased:           {CodeReceived, CodeRecalled},
	CodeReceived:            {CodeInspection, CodeQuarantined, CodeRecalled},
	CodeInspection:          {CodeStored, CodeQuarantined, CodeRecalled},
	CodeStored:              {CodeIssued, CodeTransferred, CodeExpired, CodeQuarantined, CodeRecalled},
	CodeIssued:              {CodeInUse, CodeReturned, CodeRecalled},
	CodeInUse:               {CodeTransferred, CodeReturned, CodeWasteClassification, CodeRecalled},
	CodeTransferred:         {CodeStored, CodeInUse, CodeRecalled},
	CodeReturned:            {CodeStored, CodeWasteClassification, CodeRecalled},
	CodeWasteClassification: {CodeWasteProcessing},
	CodeWasteProcessing:     {CodeDisposed},
	CodeDisposed:            {CodeClosed},
	CodeClosed:              {},
	CodeRejected:            {},
	CodeExpired:             {CodeWasteClassification},
	CodeQuarantined:         {CodeStored, CodeWasteClassification},
	CodeRecalled:            {CodeWasteClassification},
}

// ValidateTransition checks if from → to state transition is allowed in Chemical lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current Chemical state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid Chemical state transition from %s to %s", from, to)
}
