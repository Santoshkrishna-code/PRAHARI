package finding

import "time"

// Finding represents an audit finding, inspection deviation, or compliance non-conformity.
type Finding struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	SourceType  string    `json:"source_type"` // AUDIT, INSPECTION, COMPLIANCE
	SourceID    string    `json:"source_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Severity    string    `json:"severity"` // MAJOR, MINOR, OPPORTUNITY
	CreatedAt   time.Time `json:"created_at"`
}
