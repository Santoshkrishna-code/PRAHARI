package errors

import (
	"errors"
)

// Is checks if target matches err in the unwrapping chain.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's chain matching target.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Cause recursively unwraps the error chain to identify the root error trigger.
func Cause(err error) error {
	for err != nil {
		unwrapper, ok := err.(interface{ Unwrap() error })
		if !ok {
			break
		}
		next := unwrapper.Unwrap()
		if next == nil {
			break
		}
		err = next
	}
	return err
}

// GetCode extracts standard ErrorCode from error. Defaults to CodeUnknown.
func GetCode(err error) ErrorCode {
	if err == nil {
		return ""
	}
	
	var appErr *AppError
	if As(err, &appErr) {
		return appErr.Code
	}
	
	return CodeUnknown
}
