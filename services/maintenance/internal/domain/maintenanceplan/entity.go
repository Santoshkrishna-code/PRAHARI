package maintenanceplan

import (
	"errors"
	"time"
)

// Interval defines schedule cycles.
type Interval string

const (
	IntervalDaily   Interval = "DAILY"
	IntervalWeekly  Interval = "WEEKLY"
	IntervalMonthly Interval = "MONTHLY"
)

// MaintenancePlan defines pre-scheduled recurring cycles rules templates.
type MaintenancePlan struct {
	ID          string    `json:"id" db:"id"`
	AssetID     string    `json:"asset_id" db:"asset_id"`
	Title       string    `json:"title" db:"title"`
	Interval    Interval  `json:"interval" db:"interval"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	LastRunDate time.Time `json:"last_run_date" db:"last_run_date"`
	NextRunDate time.Time `json:"next_run_date" db:"next_run_date"`
}

// Validate checks domain invariants.
func (p *MaintenancePlan) Validate() error {
	if p.AssetID == "" {
		return errors.New("asset ID is required for maintenance plan")
	}
	if p.Title == "" {
		return errors.New("plan title is required")
	}
	return nil
}

// CalculateNextRun advances planned triggers dates.
func (p *MaintenancePlan) CalculateNextRun() {
	p.LastRunDate = time.Now()
	switch p.Interval {
	case IntervalDaily:
		p.NextRunDate = p.LastRunDate.AddDate(0, 0, 1)
	case IntervalWeekly:
		p.NextRunDate = p.LastRunDate.AddDate(0, 0, 7)
	case IntervalMonthly:
		p.NextRunDate = p.LastRunDate.AddDate(0, 1, 0)
	}
}
