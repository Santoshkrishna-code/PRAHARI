package sustainabilityreport

import (
	"errors"
	"time"
)

// Report tracks corporate governance reports.
type Report struct {
	ID             string    `json:"id" db:"id"`
	BusinessUnitID string    `json:"business_unit_id" db:"business_unit_id"`
	Title          string    `json:"title" db:"title"`
	ReportingYear  int       `json:"reporting_year" db:"reporting_year"`
	FrameworksUsed string    `json:"frameworks_used" db:"frameworks_used"` // e.g. "GRI,SASB"
	Status         string    `json:"status" db:"status"` // "OBJECTIVE_DEFINED", "VALIDATION", "PUBLISHED"
	ApprovedBy     string    `json:"approved_by" db:"approved_by"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks report fields.
func (r *Report) Validate() error {
	if r.BusinessUnitID == "" {
		return errors.New("business unit reference is required")
	}
	if r.Title == "" {
		return errors.New("report title description is required")
	}
	return nil
}
