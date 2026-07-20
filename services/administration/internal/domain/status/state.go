package status

// Code represents organization lifecycle state.
type Code string

const (
	CodeDraft        Code = "DRAFT"
	CodeConfigured   Code = "CONFIGURED"
	CodeValidated    Code = "VALIDATED"
	CodeActivated    Code = "ACTIVATED"
	CodeOperational  Code = "OPERATIONAL"
	CodeSuspended    Code = "SUSPENDED"
	CodeArchived     Code = "ARCHIVED"
)

func (c Code) String() string {
	return string(c)
}
