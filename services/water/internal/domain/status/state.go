package status

// Code represents water lifecycle state.
type Code string

const (
	CodeSourceRegistered  Code = "SOURCE_REGISTERED"
	CodeCollection        Code = "COLLECTION"
	CodeStorage           Code = "STORAGE"
	CodeTreatment         Code = "TREATMENT"
	CodeDistribution      Code = "DISTRIBUTION"
	CodeConsumption       Code = "CONSUMPTION"
	CodeRecyclingReuse    Code = "RECYCLING_REUSE"
	CodePerformanceReview Code = "PERFORMANCE_REVIEW"
	CodeArchived          Code = "ARCHIVED"
)

func (c Code) String() string {
	return string(c)
}
