package status

// Code represents MOC lifecycle state.
type Code string

const (
	CodeDraft            Code = "DRAFT"
	CodeImpactAssessment Code = "IMPACT_ASSESSMENT"
	CodeTechnicalReview  Code = "TECHNICAL_REVIEW"
	CodeRiskReview       Code = "RISK_REVIEW"
	CodeSafetyReview     Code = "SAFETY_REVIEW"
	CodeApproval         Code = "APPROVAL"
	CodeImplementation   Code = "IMPLEMENTATION"
	CodeVerification     Code = "VERIFICATION"
	CodeCloseout         Code = "CLOSEOUT"
	CodeRejected         Code = "REJECTED"
	CodeCancelled        Code = "CANCELLED"
	CodeRolledBack       Code = "ROLLED_BACK"
)

func (c Code) String() string {
	return string(c)
}
