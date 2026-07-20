package status

// Code represents Business Continuity Management lifecycle state.
type Code string

const (
	CodePlanned                Code = "PLANNED"
	CodeBusinessImpactAnalysis Code = "BUSINESS_IMPACT_ANALYSIS"
	CodeStrategyDevelopment    Code = "STRATEGY_DEVELOPMENT"
	CodePlanDevelopment        Code = "PLAN_DEVELOPMENT"
	CodeApproval               Code = "APPROVAL"
	CodeExercise               Code = "EXERCISE"
	CodeActivation             Code = "ACTIVATION"
	CodeRecovery               Code = "RECOVERY"
	CodeReview                 Code = "REVIEW"
	CodeContinuousImprovement  Code = "CONTINUOUS_IMPROVEMENT"
	CodeArchived               Code = "ARCHIVED"
)

func (c Code) String() string {
	return string(c)
}
