package status

// Code represents IP camera connection status.
type Code string

const (
	CodeOnline       Code = "ONLINE"
	CodeOffline      Code = "OFFLINE"
	CodeUnreachable  Code = "UNREACHABLE"
)

func (c Code) String() string {
	return string(c)
}
