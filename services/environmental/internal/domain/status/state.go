package status

type Code string

const (
	CodePlanned              Code = "PLANNED"
	CodeMonitoring           Code = "MONITORING"
	CodeSampling             Code = "SAMPLING"
	CodeLaboratoryAnalysis   Code = "LABORATORY_ANALYSIS"
	CodeComplianceEvaluation Code = "COMPLIANCE_EVALUATION"
	CodeCorrectiveAction     Code = "CORRECTIVE_ACTION"
	CodeVerification         Code = "VERIFICATION"
	CodeClosed               Code = "CLOSED"
	CodeNonCompliant         Code = "NON_COMPLIANT"
	CodeEscalated            Code = "ESCALATED"
)
