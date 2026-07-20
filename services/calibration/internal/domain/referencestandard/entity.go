package referencestandard

import "time"

// Standard represents a high-accuracy reference instrument used to calibrate field devices (provides NIST or national lab traceability).
type Standard struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	ModelName      string    `json:"model_name"`
	SerialNumber   string    `json:"serial_number"`
	LastCalibrated time.Time `json:"last_calibrated"`
	ExpiryDate     time.Time `json:"expiry_date"`
	AccuracyClass  string    `json:"accuracy_class"`
}
