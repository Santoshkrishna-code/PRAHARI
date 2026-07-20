package inspection

import "time"

// Record logs periodic safety inspections of chemical containers and storage areas.
type Record struct {
	ID            string    `json:"id"`
	StorageAreaID string    `json:"storage_area_id"`
	InspectedBy   string    `json:"inspected_by"`
	InspectedAt   time.Time `json:"inspected_at"`
	HasLeaks      bool      `json:"has_leaks"`
	LabelsIntact  bool      `json:"labels_intact"`
	StorageOk     bool      `json:"storage_ok"`
	Notes         string    `json:"notes,omitempty"`
}
