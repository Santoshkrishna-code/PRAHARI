package status

// Code represents PPE lifecycle state.
type Code string

const (
	CodeAvailable   Code = "AVAILABLE"
	CodeReserved    Code = "RESERVED"
	CodeIssued      Code = "ISSUED"
	CodeInUse       Code = "IN_USE"
	CodeInspection  Code = "INSPECTION"
	CodeReturned    Code = "RETURNED"
	CodeMaintenance Code = "MAINTENANCE"
	CodeExpired     Code = "EXPIRED"
	CodeDisposed    Code = "DISPOSED"
)

func (c Code) String() string {
	return string(c)
}
