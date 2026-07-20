package status

// Code represents PHA lifecycle state.
type Code string

const (
	CodeDraft                  Code = "DRAFT"
	CodePreparation            Code = "PREPARATION"
	CodeStudy                  Code = "STUDY"
	CodeRiskEvaluation         Code = "RISK_EVALUATION"
	CodeRecommendation         Code = "RECOMMENDATION"
	CodeApproval               Code = "APPROVAL"
	CodeImplementationTracking Code = "IMPLEMENTATION_TRACKING"
	CodeVerification           Code = "VERIFICATION"
	CodeRevalidationScheduled  Code = "REVALIDATION_SCHEDULED"
	CodeClosed                 Code = "CLOSED"
	CodeCancelled              Code = "CANCELLED"
	CodeSuperseded             Code = "SUPERSEDED"
)

func (c Code) String() string {
	return string(c)
}
