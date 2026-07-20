package ppereturn

import "time"

// Record tracks physical return of issued protective equipment, indicating wear & tear checks.
type Record struct {
	ID          string    `json:"id"`
	IssueID     string    `json:"issue_id"`
	ItemID      string    `json:"item_id"`
	ReturnedBy  string    `json:"returned_by"`
	ReturnedAt  time.Time `json:"returned_at"`
	Condition   string    `json:"condition"` // GOOD, WORN, DAMAGED
}
