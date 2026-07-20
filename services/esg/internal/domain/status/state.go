package status

type Code string

const (
	CodeObjectiveDefined     Code = "OBJECTIVE_DEFINED"
	CodeBaselineEstablished  Code = "BASELINE_ESTABLISHED"
	CodeDataCollection       Code = "DATA_COLLECTION"
	CodeCalculation          Code = "CALCULATION"
	CodeValidation           Code = "VALIDATION"
	CodeDisclosure           Code = "DISCLOSURE"
	CodeExecutiveReview      Code = "EXECUTIVE_REVIEW"
	CodePublished            Code = "PUBLISHED"
	CodeReopened             Code = "REOPENED"
	CodeArchived             Code = "ARCHIVED"
)
