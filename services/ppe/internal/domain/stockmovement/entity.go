package stockmovement

import "time"

// Transaction logs stock movements (e.g. procurement, adjustment, scrap, issue, return).
type Transaction struct {
	ID             string    `json:"id"`
	PPEID          string    `json:"ppe_id"`
	MovementType   string    `json:"movement_type"` // PROCUREMENT, SCRAPPED, ADJUSTMENT, ISSUED, RETURNED
	QuantityChange int       `json:"quantity_change"`
	RecordedBy     string    `json:"recorded_by"`
	RecordedAt     time.Time `json:"recorded_at"`
	Remarks        string    `json:"remarks"`
}
