package disposal

import "time"

// Record logs chemical waste disposal transactions.
type Record struct {
	ID           string    `json:"id"`
	ContainerID  string    `json:"container_id"`
	QtyDisposed  float64   `json:"qty_disposed"`
	UnitOfMeasure string   `json:"unit_of_measure"`
	DisposedBy   string    `json:"disposed_by"`
	DisposedAt   time.Time `json:"disposed_at"`
	FacilityName string    `json:"facility_name"` // External waste management facility
	ManifestNum  string    `json:"manifest_num"`  // EPA waste manifest tracking number
}
