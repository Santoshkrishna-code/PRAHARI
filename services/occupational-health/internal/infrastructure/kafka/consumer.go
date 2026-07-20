package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"

	clearanceApp "prahari/services/occupational-health/internal/application/clearance"
)

type PermitRequestedEvent struct {
	PermitID     string `json:"permit_id"`
	WorkerID     string `json:"worker_id"`
	PermitType   string `json:"permit_type"`
	DepartmentID string `json:"department_id"`
}

type IncidentClosedEvent struct {
	IncidentID   string `json:"incident_id"`
	DepartmentID string `json:"department_id"`
}

type HealthTrigger interface {
	TransitionStatus(ctx context.Context, cmd clearanceApp.TransitionStatusCommand) error
}

type Consumer struct {
	trigger HealthTrigger
}

func NewConsumer(trigger HealthTrigger) *Consumer {
	return &Consumer{trigger: trigger}
}

func (c *Consumer) HandlePermitRequested(ctx context.Context, data []byte) error {
	var event PermitRequestedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed permit requested event, verifying medical clearance rules",
		prahariLogger.String("permit_id", event.PermitID),
		prahariLogger.String("worker_id", event.WorkerID))

	return nil
}

func (c *Consumer) HandleIncidentClosed(ctx context.Context, data []byte) error {
	var event IncidentClosedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed incident closed event, evaluating need for health surveillance updates",
		prahariLogger.String("incident_id", event.IncidentID))

	return nil
}
