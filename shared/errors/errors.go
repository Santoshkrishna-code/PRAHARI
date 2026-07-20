package errors

import (
	"fmt"
	"io"
)

// AppError is the standard, structured error type used throughout the PRAHARI ecosystem.
type AppError struct {
	Code       ErrorCode              `json:"code"`
	Message    string                 `json:"message"`
	Err        error                  `json:"-"` // Underlying cause error
	Stack      StackTrace             `json:"stack_trace,omitempty"`
	IsRetry    bool                   `json:"retryable"`
	Diagnostics map[string]interface{} `json:"diagnostics,omitempty"`
}

// Error implements the native Go error interface.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the underlying error, implementing native Go error unwrapping support.
func (e *AppError) Unwrap() error {
	return e.Err
}

// Format overrides default print formats, enabling %+v to output stack traces.
func (e *AppError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, e.Error())
			if len(e.Stack) > 0 {
				_, _ = io.WriteString(s, "\nStack Trace:\n")
				_, _ = io.WriteString(s, e.Stack.String())
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = io.WriteString(s, fmt.Sprintf("%q", e.Error()))
	}
}

// New constructs a base AppError without wrapping an underlying error.
func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Stack:   CaptureStackTrace(1),
	}
}

// Wrap wraps an existing error into an AppError with custom context code and message.
func Wrap(err error, code ErrorCode, message string) *AppError {
	if err == nil {
		return nil
	}
	
	// If it is already an AppError, we wrap it but keep its nested diagnostics/stack if preferred
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
		Stack:   CaptureStackTrace(1),
	}
}
