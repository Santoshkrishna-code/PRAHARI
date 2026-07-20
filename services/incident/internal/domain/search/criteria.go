package search

import (
	"time"
)

// Criteria defines the composite filter parameters for searching incidents.
// All fields are optional — only non-zero values are applied as filters.
type Criteria struct {
	IncidentNumber string    `json:"incident_number,omitempty"`
	Category       string    `json:"category,omitempty"`
	Severity       string    `json:"severity,omitempty"`
	Status         string    `json:"status,omitempty"`
	Type           string    `json:"type,omitempty"`
	DateFrom       time.Time `json:"date_from,omitempty"`
	DateTo         time.Time `json:"date_to,omitempty"`
	ReporterID     string    `json:"reporter_id,omitempty"`
	AssigneeID     string    `json:"assignee_id,omitempty"`
	Department     string    `json:"department,omitempty"`
	Location       string    `json:"location,omitempty"`
	Tags           []string  `json:"tags,omitempty"`
	FreeText       string    `json:"free_text,omitempty"`
	Page           int       `json:"page,omitempty"`
	PageSize       int       `json:"page_size,omitempty"`
	SortBy         string    `json:"sort_by,omitempty"`
	SortOrder      string    `json:"sort_order,omitempty"`
}

// Result wraps paginated search results with metadata.
type Result struct {
	Items      interface{} `json:"items"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// Normalize sets default values for pagination parameters if not provided.
func (c *Criteria) Normalize() {
	if c.Page <= 0 {
		c.Page = 1
	}
	if c.PageSize <= 0 {
		c.PageSize = 20
	}
	if c.PageSize > 100 {
		c.PageSize = 100
	}
	if c.SortBy == "" {
		c.SortBy = "created_at"
	}
	if c.SortOrder == "" {
		c.SortOrder = "DESC"
	}
}

// Offset calculates the SQL offset from the page and page size.
func (c *Criteria) Offset() int {
	return (c.Page - 1) * c.PageSize
}

// HasDateRange returns true if both date boundaries are set.
func (c *Criteria) HasDateRange() bool {
	return !c.DateFrom.IsZero() && !c.DateTo.IsZero()
}

// HasFreeText returns true if a free-text search query is provided.
func (c *Criteria) HasFreeText() bool {
	return c.FreeText != ""
}
