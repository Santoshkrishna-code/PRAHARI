package errors

import (
	"net/http"
)

// HTTPStatus returns the HTTP status code matching the standard ErrorCode.
func HTTPStatus(code ErrorCode) int {
	switch code {
	case CodeInvalidArgument:
		return http.StatusBadRequest
	case CodeUnauthenticated:
		return http.StatusUnauthorized
	case CodePermissionDenied:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeAlreadyExists:
		return http.StatusConflict
	case CodeConflict:
		return http.StatusConflict
	case CodeResourceExhausted:
		return http.StatusTooManyRequests
	case CodeUnavailable:
		return http.StatusServiceUnavailable
	case CodeDeadlineExceeded:
		return http.StatusGatewayTimeout
	case CodeInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// ErrorHTTPStatus extracts ErrorCode from error and resolves corresponding status code.
func ErrorHTTPStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}
	return HTTPStatus(GetCode(err))
}
