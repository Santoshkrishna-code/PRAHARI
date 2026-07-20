package events

import (
	"context"

	prahariLogger "prahari/shared/logger"

	kafkaInfra "prahari/services/permit/internal/infrastructure/kafka"
)

// Listener hooks Kafka topic subscriptions.
type Listener struct {
	consumer *kafkaInfra.Consumer
}

// NewListener instantiates Listener.
func NewListener(consumer *kafkaInfra.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

// Start registers consume handlers.
func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Registering Permit Service Kafka listeners: workflow.completed, workflow.failed, approval.completed")
}

// Stop releases listeners.
func (l *Listener) Stop(ctx context.Context) {
	prahariLogger.Info(ctx, "Stopping Kafka event listeners")
}
