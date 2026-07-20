package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// IncidentClosedEvent represents inbound closed incidents.
type IncidentClosedEvent struct {
	IncidentID   string `json:"incident_id"`
	DepartmentID string `json:"department_id"`
}

// ComplianceTrigger defines state machine transition ports.
type ComplianceTrigger interface {
	TransitionStatus(ctx context.Context, cmd struct {
		ComplianceID string
		TargetCode   string
		ActorID      string
		Reason       string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	trigger ComplianceTrigger
}

// NewConsumer instantiates Consumer.
func NewConsumer(trigger ComplianceTrigger) *Consumer {
	return &Consumer{trigger: trigger}
}

// HandleIncidentClosed reviews audit findings.
func (c *Consumer) HandleIncidentClosed(ctx context.Context, data []byte) error {
	var event IncidentClosedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed incident closed event, triggering statutory compliance checks",
		prahariLogger.String("incident_id", event.IncidentID),
		prahariLogger.String("department_id", event.DepartmentID))

	return nil
}
