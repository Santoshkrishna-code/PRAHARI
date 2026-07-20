package investigation

import (
	"errors"
	"fmt"
	"time"
)

// Methodology defines the investigation analysis technique applied.
type Methodology string

const (
	MethodologyFiveWhy   Methodology = "FIVE_WHY"
	MethodologyFishbone  Methodology = "FISHBONE"
	MethodologyFTA       Methodology = "FAULT_TREE_ANALYSIS"
	MethodologyBowtie    Methodology = "BOWTIE"
	MethodologyTapRoot   Methodology = "TAPROOT"
)

// ValidMethodologies enumerates all accepted analysis techniques.
var ValidMethodologies = []Methodology{
	MethodologyFiveWhy,
	MethodologyFishbone,
	MethodologyFTA,
	MethodologyBowtie,
	MethodologyTapRoot,
}

// Status defines the progress state of an investigation.
type Status string

const (
	StatusOpen       Status = "OPEN"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
)

// Investigation represents a formal investigation into an incident.
type Investigation struct {
	ID              string      `json:"id" db:"id"`
	IncidentID      string      `json:"incident_id" db:"incident_id"`
	InvestigatorID  string      `json:"investigator_id" db:"investigator_id"`
	Methodology     Methodology `json:"methodology" db:"methodology"`
	Findings        string      `json:"findings,omitempty" db:"findings"`
	Recommendations string      `json:"recommendations,omitempty" db:"recommendations"`
	Status          Status      `json:"status" db:"status"`
	StartedAt       time.Time   `json:"started_at" db:"started_at"`
	CompletedAt     *time.Time  `json:"completed_at,omitempty" db:"completed_at"`
	CreatedAt       time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at" db:"updated_at"`
}

// Validate enforces domain invariants on the investigation aggregate.
func (inv *Investigation) Validate() error {
	if inv.IncidentID == "" {
		return errors.New("incident ID is required for investigation")
	}
	if inv.InvestigatorID == "" {
		return errors.New("investigator ID is required")
	}
	if !inv.isValidMethodology() {
		return fmt.Errorf("invalid investigation methodology: %s", inv.Methodology)
	}
	return nil
}

// RecordFindings captures the investigation findings and recommendations.
func (inv *Investigation) RecordFindings(findings, recommendations string) {
	inv.Findings = findings
	inv.Recommendations = recommendations
	inv.UpdatedAt = time.Now()
}

// Complete marks the investigation as completed.
func (inv *Investigation) Complete() {
	now := time.Now()
	inv.CompletedAt = &now
	inv.Status = StatusCompleted
	inv.UpdatedAt = now
}

// IsComplete returns true if the investigation has been completed.
func (inv *Investigation) IsComplete() bool {
	return inv.Status == StatusCompleted
}

// isValidMethodology checks whether the methodology is among accepted techniques.
func (inv *Investigation) isValidMethodology() bool {
	for _, m := range ValidMethodologies {
		if inv.Methodology == m {
			return true
		}
	}
	return false
}
