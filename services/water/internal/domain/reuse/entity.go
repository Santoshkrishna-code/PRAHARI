package reuse

import "time"

// Program represents a water reuse initiative.
type Program struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	ProgramName     string    `json:"program_name"`
	ReuseApplication string  `json:"reuse_application"`
	SourceStream    string    `json:"source_stream"`
	VolumeKLD       float64   `json:"volume_kld"`
	QualityGrade    string    `json:"quality_grade"`
	Status          string    `json:"status"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
