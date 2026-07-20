package events

import "time"

const (
	EventChemicalCreated           = "chemical.created"
	EventChemicalApproved          = "chemical.approved"
	EventChemicalReceived          = "chemical.received"
	EventChemicalStored            = "chemical.stored"
	EventChemicalIssued            = "chemical.issued"
	EventChemicalTransferred       = "chemical.transferred"
	EventChemicalExpired           = "chemical.expired"
	EventChemicalRecalled          = "chemical.recalled"
	EventChemicalDisposed          = "chemical.disposed"
	EventChemicalThresholdExceeded = "chemical.threshold.exceeded"
	EventChemicalSpillDetected     = "chemical.spill.detected"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
