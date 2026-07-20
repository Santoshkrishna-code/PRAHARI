package prooftest

import "time"

// Test represents an IEC 61511 proof test execution for a barrier/SIF.
type Test struct {
	ID          string    `json:"id"`
	BarrierID   string    `json:"barrier_id"`
	TestNumber  string    `json:"test_number"`
	ExecutedBy  string    `json:"executed_by"`
	Passed      bool      `json:"passed"`
	AsFoundPFD  float64   `json:"as_found_pfd"`
	AsLeftPFD   float64   `json:"as_left_pfd"`
	WorkOrderID string    `json:"work_order_id,omitempty"`
	Notes       string    `json:"notes"`
	ExecutedAt  time.Time `json:"executed_at"`
}
