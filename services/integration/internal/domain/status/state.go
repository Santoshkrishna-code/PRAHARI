package status

// Code represents connector state.
type Code string

const (
	CodeDisconnected Code = "DISCONNECTED"
	CodeConnecting   Code = "CONNECTING"
	CodeConnected    Code = "CONNECTED"
	CodeFailed       Code = "FAILED"
)

func (c Code) String() string {
	return string(c)
}
