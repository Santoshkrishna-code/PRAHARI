package status

// Code represents document parsing and indexing status lifecycle.
type Code string

const (
	CodeQueued     Code = "QUEUED"
	CodeParsing    Code = "PARSING"
	CodeVectorized Code = "VECTORIZED"
	CodeFailed     Code = "FAILED"
)

func (c Code) String() string {
	return string(c)
}
