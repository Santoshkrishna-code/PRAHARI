package safetyinstrumentedsystem

import "time"

// SIS represents a Safety Instrumented System logic solver / hardware architecture.
type SIS struct {
	ID           string    `json:"id"`
	BarrierID    string    `json:"barrier_id"`
	SISName      string    `json:"sis_name"`
	Architecture string    `json:"architecture"` // 1oo1, 1oo2, 2oo3
	Vendor       string    `json:"vendor"`
	LogicSolver  string    `json:"logic_solver"`
	CreatedAt    time.Time `json:"created_at"`
}
