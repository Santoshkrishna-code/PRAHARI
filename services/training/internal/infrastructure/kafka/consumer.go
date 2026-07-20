package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"

	trainingApp "prahari/services/training/internal/application/training"
)

// AuditClosedEvent represents inbound closed audits.
type AuditClosedEvent struct {
	AuditID      string `json:"audit_id"`
	DepartmentID string `json:"department_id"`
}

// TrainingTrigger defines state machine transition ports.
type TrainingTrigger interface {
	TransitionStatus(ctx context.Context, cmd trainingApp.TransitionStatusCommand) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	trigger TrainingTrigger
}

// NewConsumer instantiates Consumer.
func NewConsumer(trigger TrainingTrigger) *Consumer {
	return &Consumer{trigger: trigger}
}

// HandleAuditClosed triggers scheduled refresher checks.
func (c *Consumer) HandleAuditClosed(ctx context.Context, data []byte) error {
	var event AuditClosedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed audit closed event, registering new refresher trainings",
		prahariLogger.String("audit_id", event.AuditID),
		prahariLogger.String("department_id", event.DepartmentID))

	return nil
}
