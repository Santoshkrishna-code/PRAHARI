package warranty

import (
	"errors"
	"time"
)

// Warranty maps maintenance coverage.
type Warranty struct {
	ID             string    `json:"id" db:"id"`
	AssetID        string    `json:"asset_id" db:"asset_id"`
	StartDate      time.Time `json:"start_date" db:"start_date"`
	EndDate        time.Time `json:"end_date" db:"end_date"`
	CoverageDetail string    `json:"coverage_detail" db:"coverage_detail"`
	ContactPerson  string    `json:"contact_person" db:"contact_person"`
	ContactEmail   string    `json:"contact_email" db:"contact_email"`
}

// Validate checks domain invariants.
func (w *Warranty) Validate() error {
	if w.AssetID == "" {
		return errors.New("asset ID is required for warranty")
	}
	if w.EndDate.Before(w.StartDate) {
		return errors.New("warranty end date must be after start date")
	}
	return nil
}

// IsExpired checks warranty validity windows.
func (w *Warranty) IsExpired() bool {
	return time.Now().After(w.EndDate)
}
