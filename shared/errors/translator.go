package errors

import (
	"database/sql"
	"errors"
	"net"
	"strings"
)

// TranslateDatabaseError inspects database driver exceptions and maps them to standard ErrorCodes.
func TranslateDatabaseError(err error, resourceName string) error {
	if err == nil {
		return nil
	}

	// 1. Check standard sql.ErrNoRows
	if errors.Is(err, sql.ErrNoRows) {
		return &AppError{
			Code:    CodeNotFound,
			Message: resourceName + " not found",
			Err:     err,
			Stack:   CaptureStackTrace(1),
		}
	}

	errMsg := err.Error()

	// 2. Check standard Postgres Unique Violations (Code 23505) or duplicate keys
	if strings.Contains(errMsg, "duplicate key") || strings.Contains(errMsg, "23505") {
		return &AppError{
			Code:    CodeAlreadyExists,
			Message: resourceName + " already exists",
			Err:     err,
			Stack:   CaptureStackTrace(1),
		}
	}

	// 3. Check Postgres Foreign Key Violations (Code 23503)
	if strings.Contains(errMsg, "violates foreign key") || strings.Contains(errMsg, "23503") {
		return &AppError{
			Code:    CodeConflict,
			Message: resourceName + " violates relational constraints",
			Err:     err,
			Stack:   CaptureStackTrace(1),
		}
	}

	return &AppError{
		Code:    CodeInternal,
		Message: "Database operation failed",
		Err:     err,
		Stack:   CaptureStackTrace(1),
	}
}

// TranslateNetworkError checks network timeouts and marks them transient (retryable).
func TranslateNetworkError(err error) error {
	if err == nil {
		return nil
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		return &AppError{
			Code:    CodeUnavailable,
			Message: "Network connection unavailable",
			Err:     err,
			IsRetry: netErr.Timeout() || netErr.Temporary(),
			Stack:   CaptureStackTrace(1),
		}
	}

	return &AppError{
		Code:    CodeUnknown,
		Message: "Network operation failed",
		Err:     err,
		Stack:   CaptureStackTrace(1),
	}
}
