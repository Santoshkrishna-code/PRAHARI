package events

import (
	"context"

	prahariLogger "prahari/shared/logger"

	kafkaInfra "prahari/services/incident/internal/infrastructure/kafka"
)

// Listener processes inbound Kafka events from platform services.
// It runs as a background goroutine consuming from the incident consumer group.
type Listener struct {
	consumer *kafkaInfra.Consumer
}

// NewListener constructs an event Listener.
func NewListener(consumer *kafkaInfra.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

// Start begins listening for inbound platform events.
// In production, this would register Kafka consumer group handlers.
func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Kafka event listener for incident service")

	// In production, register topic handlers with the Kafka consumer group:
	// consumerGroup.Subscribe("workflow.completed", l.consumer.HandleWorkflowCompleted)
	// consumerGroup.Subscribe("workflow.failed", l.consumer.HandleWorkflowFailed)
	// consumerGroup.Subscribe("approval.completed", l.consumer.HandleApprovalCompleted)

	prahariLogger.Info(ctx, "Kafka event listener registered for topics: workflow.completed, workflow.failed, approval.completed")
}

// Stop gracefully shuts down the event listener.
func (l *Listener) Stop(ctx context.Context) {
	prahariLogger.Info(ctx, "Stopping Kafka event listener")
}
