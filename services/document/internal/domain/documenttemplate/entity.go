package documenttemplate

import "time"

// Template represents a standardized document template (e.g., standard SOP layout).
type Template struct {
	ID           string    `json:"id"`
	TemplateName string    `json:"template_name"`
	DocumentType string    `json:"document_type"`
	FileURL      string    `json:"file_url"`
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
}
