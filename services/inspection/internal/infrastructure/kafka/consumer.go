package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	prahariLogger "prahari/shared/logger"
)

// WorkflowEvent represents inbound workflow updates message schemas.
type WorkflowEvent struct {
	WorkflowID   string `json:"workflow_id"`
	InspectionID string `json:"inspection_id"`
	Status       string `json:"status"`
}

// InspectionStatusUpdater updates inspection status based on incoming messages.
type InspectionStatusUpdater interface {
	TransitionStatus(ctx context.Context, cmd struct {
		InspectionID string
		TargetCode   string
		ActorID      string
		Reason       string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	updater InspectionStatusUpdater
}

// NewConsumer instantiates Consumer.
func NewConsumer(updater InspectionStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandleWorkflowCompleted processes completion checks.
func (c *Consumer) HandleWorkflowCompleted(ctx context.Context, data []byte) error {
	var event WorkflowEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed workflow completion event for inspection",
		prahariLogger.String("workflow_id", event.WorkflowID),
		prahariLogger.String("inspection_id", event.InspectionID))

	return nil
}
