package errors

import (
	"encoding/json"
	"net/http"
)

// NewBadRequestError constructs a new invalid argument error wrapper.
func NewBadRequestError(message string, err error) *AppError {
	return Wrap(err, CodeInvalidArgument, message)
}

// NewNotFoundError constructs a new resource not found error wrapper.
func NewNotFoundError(message string, err error) *AppError {
	return Wrap(err, CodeNotFound, message)
}

// WriteHTTP serializes error payloads matching HTTP status codes.
func WriteHTTP(w http.ResponseWriter, err error) {
	status := ErrorHTTPStatus(err)
	prob := NewProblemDetails(err, status, "An error occurred", "")
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(prob)
}
