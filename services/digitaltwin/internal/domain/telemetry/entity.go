package telemetry

import "time"

// Binding maps a sensor tag to a digital twin equipment node.
type Binding struct {
	ID          string    `json:"id"`
	TwinID      string    `json:"twin_id"`
	EquipmentID string    `json:"equipment_id"`
	SensorTag   string    `json:"sensor_tag"` // E.g., TI-101, PI-202
	Unit        string    `json:"unit"` // °C, bar, m³/h
	UpdatedAt   time.Time `json:"updated_at"`
}
