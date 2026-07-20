package alternatefacility

import "time"

// Facility represents a backup control room, hot site, cold site, or secondary office location.
type Facility struct {
	ID           string    `json:"id"`
	FacilityCode string    `json:"facility_code"`
	Name         string    `json:"name"`
	FacilityType string    `json:"facility_type"` // HOT_SITE, WARM_SITE, COLD_SITE, SECONDARY_CONTROL_ROOM
	Location     string    `json:"location"`
	SeatingCap   int       `json:"seating_cap"`
	CreatedAt    time.Time `json:"created_at"`
}
