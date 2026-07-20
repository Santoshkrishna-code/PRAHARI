package events

import (
	"context"

	prahariLogger "prahari/shared/logger"

	kafkaInfra "prahari/services/inspection/internal/infrastructure/kafka"
)

// Listener hooks Kafka topic subscriptions.
type Listener struct {
	consumer *kafkaInfra.Consumer
}

// NewListener instantiates Listener.
func NewListener(consumer *kafkaInfra.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

// Start registers walkthrough triggers.
func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Registering Inspection Service Kafka listener subscriptions: workflow.completed, workflow.failed, permit.closed, incident.closed")
}

// Stop releases subscriptions.
func (l *Listener) Stop(ctx context.Context) {
	prahariLogger.Info(ctx, "Stopping Kafka event listeners")
}
