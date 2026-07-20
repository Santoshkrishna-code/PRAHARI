package search

import (
	"time"
)

// Criteria defines filter parameters for list queries.
type Criteria struct {
	HazardNumber string    `json:"hazard_number,omitempty"`
	HazardType   string    `json:"hazard_type,omitempty"`
	Category     string    `json:"category,omitempty"`
	AssetID      string    `json:"asset_id,omitempty"`
	ContractorID string    `json:"contractor_id,omitempty"`
	DepartmentID string    `json:"department_id,omitempty"`
	RiskLevel    string    `json:"risk_level,omitempty"`
	Status       string    `json:"status,omitempty"`
	OwnerID      string    `json:"owner_id,omitempty"`
	DateFrom     time.Time `json:"date_from,omitempty"`
	DateTo       time.Time `json:"date_to,omitempty"`
	FreeText     string    `json:"free_text,omitempty"`
	Page         int       `json:"page,omitempty"`
	PageSize     int       `json:"page_size,omitempty"`
	SortBy       string    `json:"sort_by,omitempty"`
	SortOrder    string    `json:"sort_order,omitempty"`
}

// Result holds return items list and page totals.
type Result struct {
	Items      interface{} `json:"items"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// Normalize sets pagination criteria fallback defaults.
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

// Offset returns DB listing skip offsets.
func (c *Criteria) Offset() int {
	return (c.Page - 1) * c.PageSize
}
