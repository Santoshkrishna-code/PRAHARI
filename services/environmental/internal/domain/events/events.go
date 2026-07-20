package events

import (
	"time"
)

// EnvironmentRecordCreated represents new aspect logs.
type EnvironmentRecordCreated struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
}

// SamplingCompleted represents sampling logs.
type SamplingCompleted struct {
	ID           string    `json:"id"`
	SampleNumber string    `json:"sample_number"`
	Timestamp    time.Time `json:"timestamp"`
}

// LabAnalysisCompleted represents lab tests logs.
type LabAnalysisCompleted struct {
	ID        string    `json:"id"`
	SampleID  string    `json:"sample_id"`
	Timestamp time.Time `json:"timestamp"`
}

// ComplianceEvaluationFailed represents violations.
type ComplianceEvaluationFailed struct {
	ID        string    `json:"id"`
	SourceType string   `json:"source_type"`
	Reason    string    `json:"reason"`
	Timestamp time.Time `json:"timestamp"`
}

// CorrectiveActionCreated represents CAPA actions.
type CorrectiveActionCreated struct {
	ID         string    `json:"id"`
	SourceID   string    `json:"source_id"`
	AssignedTo string    `json:"assigned_to"`
	Timestamp  time.Time `json:"timestamp"`
}

// SpillReported represents hazmat releases.
type SpillReported struct {
	ID           string    `json:"id"`
	ChemicalName string    `json:"chemical_name"`
	Volume       float64   `json:"volume"`
	Timestamp    time.Time `json:"timestamp"`
}
