package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	prahariLogger "prahari/shared/logger"
)

// WorkflowEvent defines payload consumed from Workflow Engine.
type WorkflowEvent struct {
	WorkflowID string `json:"workflow_id"`
	PermitID   string `json:"permit_id"`
	Status     string `json:"status"`
}

// PermitStatusUpdater updates permit status based on incoming messages.
type PermitStatusUpdater interface {
	TransitionStatus(ctx context.Context, permitID, status string) error
}

// Consumer handles incoming messages from Kafka.
type Consumer struct {
	updater PermitStatusUpdater
}

// NewConsumer instantiates Consumer.
func NewConsumer(updater PermitStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandleWorkflowCompleted processes workflow.completed events.
func (c *Consumer) HandleWorkflowCompleted(ctx context.Context, data []byte) error {
	var event WorkflowEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed workflow.completed event",
		prahariLogger.String("workflow_id", event.WorkflowID),
		prahariLogger.String("permit_id", event.PermitID))

	return c.updater.TransitionStatus(ctx, event.PermitID, event.Status)
}

// HandleWorkflowFailed processes workflow failures.
func (c *Consumer) HandleWorkflowFailed(ctx context.Context, data []byte) error {
	var event WorkflowEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Error(ctx, "Consumed workflow.failed event",
		prahariLogger.String("workflow_id", event.WorkflowID),
		prahariLogger.String("permit_id", event.PermitID))

	return nil
}
