package sustainabilityinitiative

import (
	"errors"
	"time"
)

// Initiative tracks specific sustainability actions.
type Initiative struct {
	ID             string    `json:"id" db:"id"`
	ObjectiveID    string    `json:"objective_id" db:"objective_id"`
	Title          string    `json:"title" db:"title"` // e.g. "Install LED lighting in Plant A"
	Description    string    `json:"description" db:"description"`
	BudgetUSD      float64   `json:"budget_usd" db:"budget_usd"`
	ActualCostUSD  float64   `json:"actual_cost_usd" db:"actual_cost_usd"`
	Co2SavedKg     float64   `json:"co2_saved_kg" db:"co2_saved_kg"`
	Status         string    `json:"status" db:"status"` // "PLANNED", "IN_PROGRESS", "COMPLETED"
	StartDate      time.Time `json:"start_date" db:"start_date"`
	EndDate        time.Time `json:"end_date" db:"end_date"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks fields.
func (i *Initiative) Validate() error {
	if i.ObjectiveID == "" {
		return errors.New("parent sustainability objective ID is required")
	}
	if i.Title == "" {
		return errors.New("initiative title is required")
	}
	return nil
}
