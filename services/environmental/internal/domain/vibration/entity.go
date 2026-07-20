package vibration

import (
	"errors"
	"time"
)

// VibrationMonitoring registers mechanical vibration parameters.
type VibrationMonitoring struct {
	ID             string    `json:"id" db:"id"`
	SourceAssetID  string    `json:"source_asset_id" db:"source_asset_id"`
	FrequencyHz    float64   `json:"frequency_hz" db:"frequency_hz"`
	VelocityMms    float64   `json:"velocity_mms" db:"velocity_mms"` // peak velocity
	LimitThreshold float64   `json:"limit_threshold" db:"limit_threshold"`
	IsOverLimit     bool      `json:"is_over_limit" db:"is_over_limit"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks vibration records.
func (v *VibrationMonitoring) Validate() error {
	if v.SourceAssetID == "" {
		return errors.New("source asset ID reference is required")
	}
	return nil
}
