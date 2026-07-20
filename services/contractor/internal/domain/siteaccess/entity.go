package siteaccess

import (
	"errors"
	"time"
)

// SiteAccess tracks location gate checks entry clearances.
type SiteAccess struct {
	ID               string    `json:"id" db:"id"`
	WorkerID         string    `json:"worker_id" db:"worker_id"`
	AllowedLocations string    `json:"allowed_locations" db:"allowed_locations"`
	BadgeNumber      string    `json:"badge_number" db:"badge_number"`
	AccessStart      time.Time `json:"access_start" db:"access_start"`
	AccessEnd        time.Time `json:"access_end" db:"access_end"`
}

// Validate checks domain invariants.
func (sa *SiteAccess) Validate() error {
	if sa.WorkerID == "" {
		return errors.New("worker ID reference is required")
	}
	if sa.BadgeNumber == "" {
		return errors.New("access badge number code is required")
	}
	return nil
}
