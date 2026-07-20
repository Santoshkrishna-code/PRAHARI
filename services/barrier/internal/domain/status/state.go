package status

// Code represents barrier lifecycle state.
type Code string

const (
	CodeRegistered          Code = "REGISTERED"
	CodeAssigned            Code = "ASSIGNED"
	CodeOperational         Code = "OPERATIONAL"
	CodeInspection          Code = "INSPECTION"
	CodeProofTest           Code = "PROOF_TEST"
	CodeIntegrityAssessment Code = "INTEGRITY_ASSESSMENT"
	CodeVerified            Code = "VERIFIED"
	CodeBypassed            Code = "BYPASSED"
	CodeImpaired            Code = "IMPAIRED"
	CodeOutOfService        Code = "OUT_OF_SERVICE"
	CodeRetired             Code = "RETIRED"
)

func (c Code) String() string {
	return string(c)
}
