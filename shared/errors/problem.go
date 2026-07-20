package errors

// ValidationErrorDetail maps specific fields to validation errors.
type ValidationErrorDetail struct {
	Name    string `json:"name"`
	Reason  string `json:"reason"`
}

// ProblemDetails represents the standard RFC 7807 error schema payload.
type ProblemDetails struct {
	Type          string                  `json:"type"` // URI identifier
	Title         string                  `json:"title"`
	Status        int                     `json:"status"`
	Detail        string                  `json:"detail,omitempty"`
	Instance      string                  `json:"instance,omitempty"`
	Code          string                  `json:"code"`
	InvalidParams []ValidationErrorDetail `json:"invalid_params,omitempty"`
}

// NewProblemDetails maps any error into the RFC 7807 response schema.
func NewProblemDetails(err error, status int, title, instance string) *ProblemDetails {
	if err == nil {
		return nil
	}
	
	code := string(CodeUnknown)
	detail := err.Error()
	
	var appErr *AppError
	if As(err, &appErr) {
		code = string(appErr.Code)
		detail = appErr.Message
	}
	
	return &ProblemDetails{
		Type:     "about:blank", // Default schema
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
		Code:     code,
	}
}
