package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodePlanned:              {CodeMonitoring, CodeClosed},
	CodeMonitoring:           {CodeSampling, CodeComplianceEvaluation, CodeClosed},
	CodeSampling:             {CodeLaboratoryAnalysis, CodeComplianceEvaluation, CodeClosed},
	CodeLaboratoryAnalysis:   {CodeComplianceEvaluation, CodeClosed},
	CodeComplianceEvaluation: {CodeCorrectiveAction, CodeClosed, CodeNonCompliant, CodeEscalated},
	CodeCorrectiveAction:     {CodeVerification, CodeClosed, CodeNonCompliant, CodeEscalated},
	CodeVerification:         {CodeClosed, CodeNonCompliant, CodeEscalated},
	CodeNonCompliant:         {CodeCorrectiveAction, CodeEscalated, CodeClosed},
	CodeEscalated:            {CodeCorrectiveAction, CodeClosed},
	CodeClosed:               {CodePlanned, CodeMonitoring},
}

// ValidateTransition checks if from → to state transition is allowed.
func ValidateTransition(from, to Code) error {
	allowed, exists := transitionMatrix[from]
	if !exists {
		return fmt.Errorf("unknown environmental state: %s", from)
	}

	for _, t := range allowed {
		if t == to {
			return nil
		}
	}

	return fmt.Errorf("invalid environmental state transition: %s → %s", from, to)
}
