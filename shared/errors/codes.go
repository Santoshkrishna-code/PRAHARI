package errors

// ErrorCode defines a standard string representation of system error types.
type ErrorCode string

const (
	CodeUnknown          ErrorCode = "UNKNOWN"
	CodeInternal         ErrorCode = "INTERNAL"
	CodeInvalidArgument  ErrorCode = "INVALID_ARGUMENT"
	CodeUnauthenticated  ErrorCode = "UNAUTHENTICATED"
	CodePermissionDenied ErrorCode = "PERMISSION_DENIED"
	CodeNotFound         ErrorCode = "NOT_FOUND"
	CodeAlreadyExists    ErrorCode = "ALREADY_EXISTS"
	CodeConflict         ErrorCode = "CONFLICT"
	CodeResourceExhausted ErrorCode = "RESOURCE_EXHAUSTED"
	CodeUnavailable      ErrorCode = "UNAVAILABLE"
	CodeDeadlineExceeded ErrorCode = "DEADLINE_EXCEEDED"
)
