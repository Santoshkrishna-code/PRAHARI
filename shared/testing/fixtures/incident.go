package fixtures

// NewIncidentPayload returns standard incident payload JSON.
func NewIncidentPayload() []byte {
	return []byte(`{
		"incident_id": "inc-101",
		"status": "CRITICAL",
		"message": "Database memory threshold crossed"
	}`)
}
