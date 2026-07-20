package events

import (
	"time"
)

// ObjectiveCreated event.
type ObjectiveCreated struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
}

// CarbonCalculated event.
type CarbonCalculated struct {
	ID         string    `json:"id"`
	TotalCo2Kg float64   `json:"total_co2_kg"`
	Timestamp  time.Time `json:"timestamp"`
}

// ScorecardUpdated event.
type ScorecardUpdated struct {
	ID           string    `json:"id"`
	OverallScore float64   `json:"overall_score"`
	Timestamp    time.Time `json:"timestamp"`
}

// DisclosurePublished event.
type DisclosurePublished struct {
	ID            string    `json:"id"`
	ReferenceCode string    `json:"reference_code"`
	Timestamp     time.Time `json:"timestamp"`
}

// ReportGenerated event.
type ReportGenerated struct {
	ID            string    `json:"id"`
	ReportingYear int       `json:"reporting_year"`
	Timestamp     time.Time `json:"timestamp"`
}

// GoalAchieved event.
type GoalAchieved struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
}
