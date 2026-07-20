package overlay

import "time"

// Layer represents a visual overlay layer (AI insights, vision detections, alarms).
type Layer struct {
	ID        string    `json:"id"`
	TwinID    string    `json:"twin_id"`
	LayerType string    `json:"layer_type"` // AI_INSIGHT, VISION_EVENT, ALARM, ANALYTICS
	SourceID  string    `json:"source_id"` // Reference to originating service event
	Label     string    `json:"label"`
	Metadata  string    `json:"metadata"` // JSON overlay data
	Timestamp time.Time `json:"timestamp"`
}
