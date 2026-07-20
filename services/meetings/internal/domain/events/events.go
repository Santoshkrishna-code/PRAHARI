package events

import "time"

const (
	EventMeetingCreated      = "meeting.created"
	EventMeetingStarted      = "meeting.started"
	EventAttendanceRecorded  = "attendance.recorded"
	EventMeetingClosed       = "meeting.closed"
	EventToolboxTalkCompleted = "toolboxtalk.completed"
	EventActionsGenerated    = "actions.generated"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
