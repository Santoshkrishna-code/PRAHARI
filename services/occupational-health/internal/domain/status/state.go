package status

// Code defines the state code for the occupational health lifecycle.
type Code string

const (
	CodeScheduled          Code = "SCHEDULED"
	CodeMedicalExamination Code = "MEDICAL_EXAMINATION"
	CodeLaboratoryTesting  Code = "LABORATORY_TESTING"
	CodePhysicianReview    Code = "PHYSICIAN_REVIEW"
	CodeFitnessAssessment  Code = "FITNESS_ASSESSMENT"
	CodeMedicalClearance   Code = "MEDICAL_CLEARANCE"
	CodeActiveMonitoring   Code = "ACTIVE_MONITORING"
	CodePeriodicReview     Code = "PERIODIC_REVIEW"
	CodeRestricted         Code = "RESTRICTED"
	CodeTemporarilyUnfit   Code = "TEMPORARILY_UNFIT"
	CodePermanentlyUnfit   Code = "PERMANENTLY_UNFIT"
	CodeExpired            Code = "EXPIRED"
)
