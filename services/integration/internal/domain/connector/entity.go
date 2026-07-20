package connector

import "time"

// Connector represents a configured external connection instance (SAP, Modbus, etc.).
type Connector struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"` // SAP, MAXIMO, OPCUA, MQTT, WEBHOUSE, TEAMS
	Status      string    `json:"status"` // DISCONNECTED, CONNECTING, CONNECTED, FAILED
	Host        string    `json:"host"`
	Port        int       `json:"port"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
