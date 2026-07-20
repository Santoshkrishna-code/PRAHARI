package reliefdevice

import "time"

// Device represents a Pressure Relief Valve (PRV/PSV) or Rupture Disc.
type Device struct {
	ID             string    `json:"id"`
	BarrierID      string    `json:"barrier_id"`
	TagNumber      string    `json:"tag_number"`
	SetPressureBar float64   `json:"set_pressure_bar"`
	CapacityKgH    float64   `json:"capacity_kg_h"`
	LastTestedAt   time.Time `json:"last_tested_at"`
	CreatedAt      time.Time `json:"created_at"`
}
