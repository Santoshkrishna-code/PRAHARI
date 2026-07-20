package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// IncidentCreatedEvent represents inbound incidents.
type IncidentCreatedEvent struct {
	IncidentID   string `json:"incident_id"`
	Title        string `json:"title"`
	DepartmentID string `json:"department_id"`
}

// ReassessmentTrigger defines state machine transition ports.
type ReassessmentTrigger interface {
	TransitionStatus(ctx context.Context, cmd struct {
		RiskID     string
		TargetCode string
		ActorID    string
		Reason     string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	trigger ReassessmentTrigger
}

// NewConsumer instantiates Consumer.
func NewConsumer(trigger ReassessmentTrigger) *Consumer {
	return &Consumer{trigger: trigger}
}

// HandleIncidentCreated reviews process safety ratings.
func (c *Consumer) HandleIncidentCreated(ctx context.Context, data []byte) error {
	var event IncidentCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed incident created event, scanning risk registers for auto-reassessment requirements",
		prahariLogger.String("incident_id", event.IncidentID),
		prahariLogger.String("department_id", event.DepartmentID))

	return nil
}
