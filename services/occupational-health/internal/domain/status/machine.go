package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeScheduled:          {CodeMedicalExamination, CodeRestricted, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodeMedicalExamination: {CodeLaboratoryTesting, CodePhysicianReview, CodeFitnessAssessment, CodeRestricted, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodeLaboratoryTesting:  {CodePhysicianReview, CodeFitnessAssessment, CodeRestricted, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodePhysicianReview:    {CodeFitnessAssessment, CodeMedicalClearance, CodeRestricted, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodeFitnessAssessment:  {CodeMedicalClearance, CodeRestricted, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodeMedicalClearance:   {CodeActiveMonitoring, CodePeriodicReview, CodeRestricted, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodeActiveMonitoring:   {CodePeriodicReview, CodeScheduled, CodeMedicalExamination, CodeRestricted, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodePeriodicReview:     {CodeScheduled, CodeMedicalExamination, CodeRestricted, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodeRestricted:         {CodeScheduled, CodeMedicalClearance, CodeActiveMonitoring, CodeTemporarilyUnfit, CodePermanentlyUnfit, CodeExpired},
	CodeTemporarilyUnfit:   {CodeScheduled, CodeMedicalExamination, CodePermanentlyUnfit, CodeExpired},
	CodePermanentlyUnfit:   {CodeExpired},
	CodeExpired:            {CodeScheduled, CodeMedicalExamination},
}

// ValidateTransition checks if from → to state transition is allowed.
func ValidateTransition(from, to Code) error {
	allowed, exists := transitionMatrix[from]
	if !exists {
		return fmt.Errorf("unknown health state: %s", from)
	}

	for _, t := range allowed {
		if t == to {
			return nil
		}
	}

	return fmt.Errorf("invalid health state transition: %s → %s", from, to)
}
