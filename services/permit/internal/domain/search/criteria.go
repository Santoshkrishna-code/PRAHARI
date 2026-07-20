package search

import (
	"time"
)

// Criteria defines filter parameters for querying permits.
type Criteria struct {
	PermitNumber string    `json:"permit_number,omitempty"`
	PermitType   string    `json:"permit_type,omitempty"`
	Status       string    `json:"status,omitempty"`
	ApplicantID  string    `json:"applicant_id,omitempty"`
	ApproverID   string    `json:"approver_id,omitempty"`
	Department   string    `json:"department,omitempty"`
	Contractor   string    `json:"contractor,omitempty"`
	WorkArea     string    `json:"work_area,omitempty"`
	DateFrom     time.Time `json:"date_from,omitempty"`
	DateTo       time.Time `json:"date_to,omitempty"`
	RiskLevel    string    `json:"risk_level,omitempty"`
	FreeText     string    `json:"free_text,omitempty"`
	Page         int       `json:"page,omitempty"`
	PageSize     int       `json:"page_size,omitempty"`
	SortBy       string    `json:"sort_by,omitempty"`
	SortOrder    string    `json:"sort_order,omitempty"`
}

// Result holds list items alongside paginated metadata totals.
type Result struct {
	Items      interface{} `json:"items"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// Normalize applies defaults for query paging inputs.
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

// Offset returns database paging index.
func (c *Criteria) Offset() int {
	return (c.Page - 1) * c.PageSize
}

// HasDateRange checks if both bounds are non-zero.
func (c *Criteria) HasDateRange() bool {
	return !c.DateFrom.IsZero() && !c.DateTo.IsZero()
}
