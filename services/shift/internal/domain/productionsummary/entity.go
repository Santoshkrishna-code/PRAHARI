package productionsummary

import "time"

// Summary compiles daily production outputs, downtime, and material flows.
type Summary struct {
	ID             string    `json:"id"`
	ShiftID        string    `json:"shift_id"`
	OutputQuantity float64   `json:"output_quantity"`
	UnitOfMeasure  string    `json:"unit_of_measure"` // Barrels, Tons, etc.
	DowntimeHours  float64   `json:"downtime_hours"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
}
