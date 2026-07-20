package protocol

// Config represents custom communication parameters (Modbus/OPC registers offset etc.).
type Config struct {
	ID          string `json:"id"`
	ConnectorID string `json:"connector_id"`
	PayloadType string `json:"payload_type"` // TEXT, RAW, COMPRESSED
	TimeoutMs   int    `json:"timeout_ms"`
}
