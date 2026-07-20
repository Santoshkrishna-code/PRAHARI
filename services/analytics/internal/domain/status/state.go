package status

// Code represents report workflow lifecycle status.
type Code string

const (
	CodeDraft     Code = "DRAFT"
	CodeGenerated Code = "GENERATED"
	CodeDelivered Code = "DELIVERED"
	CodeFailed    Code = "FAILED"
)

func (c Code) String() string {
	return string(c)
}
