package status

// Code represents the lifecycle status of a digital twin.
type Code string

const (
	CodeDraft    Code = "DRAFT"
	CodeActive   Code = "ACTIVE"
	CodeArchived Code = "ARCHIVED"
)

func (c Code) String() string {
	return string(c)
}
