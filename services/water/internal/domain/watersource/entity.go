package watersource

import "time"

// SourceType defines the origin classification.
type SourceType string

const (
	SourceMunicipal  SourceType = "MUNICIPAL"
	SourceBorewell   SourceType = "BOREWELL"
	SourceRiver      SourceType = "RIVER"
	SourceLake       SourceType = "LAKE"
	SourceReservoir  SourceType = "RESERVOIR"
	SourceRainwater  SourceType = "RAINWATER_HARVESTING"
	SourceRecycled   SourceType = "RECYCLED"
	SourceDesalinated SourceType = "DESALINATED"
)

// Source represents an enterprise water intake source.
type Source struct {
	ID           string     `json:"id"`
	PlantID      string     `json:"plant_id"`
	SourceName   string     `json:"source_name"`
	SourceType   SourceType `json:"source_type"`
	LocationCode string     `json:"location_code"`
	CapacityKLD  float64    `json:"capacity_kld"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
}
