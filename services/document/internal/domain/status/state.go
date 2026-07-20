package status

// Code represents Document Management lifecycle state.
type Code string

const (
	CodeDraft                  Code = "DRAFT"
	CodeReview                 Code = "REVIEW"
	CodeApproval               Code = "APPROVAL"
	CodePublished              Code = "PUBLISHED"
	CodeControlledDistribution Code = "CONTROLLED_DISTRIBUTION"
	CodePeriodicReview         Code = "PERIODIC_REVIEW"
	CodeRevision               Code = "REVISION"
	CodeSuperseded             Code = "SUPERSEDED"
	CodeArchived               Code = "ARCHIVED"
	CodeRejected               Code = "REJECTED"
	CodeWithdrawn              Code = "WITHDRAWN"
)

func (c Code) String() string {
	return string(c)
}
