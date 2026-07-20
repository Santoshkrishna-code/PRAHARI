package zone

import "time"

// SpatialZone represents a mapped spatial zone within a facility.
type SpatialZone struct {
	ID          string    `json:"id"`
	TwinID      string    `json:"twin_id"`
	FacilityID  string    `json:"facility_id"`
	Name        string    `json:"name"` // E.g., Hazardous Storage Area B
	ZoneType    string    `json:"zone_type"` // RESTRICTED, HAZARDOUS, SAFE_ASSEMBLY, LOADING
	Coordinates string    `json:"coordinates"` // GeoJSON polygon
	UpdatedAt   time.Time `json:"updated_at"`
}
