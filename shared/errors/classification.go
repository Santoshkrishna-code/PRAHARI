package errors

// IsRetryable returns whether the error represents a transient failure that can be retried.
func IsRetryable(err error) bool {
	if err == nil {
		return false
	}
	
	var appErr *AppError
	if As(err, &appErr) {
		return appErr.IsRetry
	}
	
	// Default: standard errors are treated as non-retryable unless classified
	return false
}

// MarkRetryable sets the retryable flag on an error if it's an AppError.
func MarkRetryable(err error) error {
	if err == nil {
		return nil
	}
	
	var appErr *AppError
	if As(err, &appErr) {
		appErr.IsRetry = true
		return appErr
	}
	
	// If not AppError, wrap it and mark retryable
	return &AppError{
		Code:    CodeUnknown,
		Message: err.Error(),
		Err:     err,
		IsRetry: true,
		Stack:   CaptureStackTrace(1),
	}
}

// MarkPermanent explicitly marks an error as permanent (non-retryable).
func MarkPermanent(err error) error {
	if err == nil {
		return nil
	}
	
	var appErr *AppError
	if As(err, &appErr) {
		appErr.IsRetry = false
		return appErr
	}
	
	return &AppError{
		Code:    CodeUnknown,
		Message: err.Error(),
		Err:     err,
		IsRetry: false,
		Stack:   CaptureStackTrace(1),
	}
}
