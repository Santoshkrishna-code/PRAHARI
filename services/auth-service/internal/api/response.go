package api

import (
	"encoding/json"
	"net/http"
)

// ResponseEnvelope represents the standard success response wrapper.
type ResponseEnvelope struct {
	Data interface{} `json:"data,omitempty"`
}

// ErrorDetail represents specific fields validation issue.
type ErrorDetail struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}

// ErrorBlock defines the standard REST error body payload.
type ErrorBlock struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details,omitempty"`
}

// ErrorEnvelope wraps the standard REST error payload.
type ErrorEnvelope struct {
	Error ErrorBlock `json:"error"`
}

// WriteJSON sends a successful JSON response payload.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// WriteError sends a standard JSON error response payload.
func WriteError(w http.ResponseWriter, status int, code, message string, details []ErrorDetail) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	envelope := ErrorEnvelope{
		Error: ErrorBlock{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
	_ = json.NewEncoder(w).Encode(envelope)
}
