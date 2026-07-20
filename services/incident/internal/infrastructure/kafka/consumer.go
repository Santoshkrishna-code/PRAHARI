package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	prahariLogger "prahari/shared/logger"
)

// WorkflowEvent represents an inbound event from the Workflow Engine.
type WorkflowEvent struct {
	WorkflowID string `json:"workflow_id"`
	IncidentID string `json:"incident_id"`
	Status     string `json:"status"`
	StepID     string `json:"step_id"`
}

// IncidentStatusUpdater defines the port for updating incident status from consumed events.
type IncidentStatusUpdater interface {
	TransitionFromWorkflow(ctx context.Context, incidentID, workflowStatus string) error
}

// Consumer processes inbound Kafka events from platform services.
type Consumer struct {
	updater IncidentStatusUpdater
}

// NewConsumer constructs a Consumer.
func NewConsumer(updater IncidentStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandleWorkflowCompleted processes workflow.completed events,
// auto-advancing incident status based on workflow outcome.
func (c *Consumer) HandleWorkflowCompleted(ctx context.Context, data []byte) error {
	var event WorkflowEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("kafka consumer: failed to deserialize workflow event: %w", err)
	}

	prahariLogger.Info(ctx, "Processing workflow.completed event",
		prahariLogger.String("workflow_id", event.WorkflowID),
		prahariLogger.String("incident_id", event.IncidentID))

	return c.updater.TransitionFromWorkflow(ctx, event.IncidentID, event.Status)
}

// HandleWorkflowFailed processes workflow.failed events.
func (c *Consumer) HandleWorkflowFailed(ctx context.Context, data []byte) error {
	var event WorkflowEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("kafka consumer: failed to deserialize workflow failure: %w", err)
	}

	prahariLogger.Error(ctx, "Workflow failed for incident",
		prahariLogger.String("workflow_id", event.WorkflowID),
		prahariLogger.String("incident_id", event.IncidentID))

	return nil
}

// HandleApprovalCompleted processes approval.completed events,
// transitioning incidents from UnderReview to Assigned.
func (c *Consumer) HandleApprovalCompleted(ctx context.Context, data []byte) error {
	var event WorkflowEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("kafka consumer: failed to deserialize approval event: %w", err)
	}

	prahariLogger.Info(ctx, "Processing approval.completed event",
		prahariLogger.String("incident_id", event.IncidentID))

	return c.updater.TransitionFromWorkflow(ctx, event.IncidentID, "APPROVED")
}
