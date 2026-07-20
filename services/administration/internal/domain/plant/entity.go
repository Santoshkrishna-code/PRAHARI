package plant

import "time"

// Plant represents a specific manufacturing, process, or production facility.
type Plant struct {
	ID             string    `json:"id"`
	BusinessUnitID string    `json:"business_unit_id"`
	Name           string    `json:"name"`
	Code           string    `json:"code"`
	Location       string    `json:"location"`
	TimeZone       string    `json:"time_zone"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
