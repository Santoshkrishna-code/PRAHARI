package endpoint

// Endpoint represents a resource path or topic definition under a connector.
type Endpoint struct {
	ID          string `json:"id"`
	ConnectorID string `json:"connector_id"`
	Path        string `json:"path"`   // E.g., /v1/assets/sync, telemetry/gas-detector
	Method      string `json:"method"` // GET, POST, SUBSCRIBE, PUBLISH
}
