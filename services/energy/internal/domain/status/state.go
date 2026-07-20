package status

type Code string

const (
	CodeMeterRegistered Code = "METER_REGISTERED"
	CodeDataCollection  Code = "DATA_COLLECTION"
	CodeValidation      Code = "VALIDATION"
	CodeAggregation     Code = "AGGREGATION"
	CodeAnalysis        Code = "ANALYSIS"
	CodeOptimization    Code = "OPTIMIZATION"
	CodeReporting       Code = "REPORTING"
	CodeArchived        Code = "ARCHIVED"
)
