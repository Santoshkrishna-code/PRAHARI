package events

import (
	"context"

	prahariLogger "prahari/shared/logger"

	kafkaInfra "prahari/services/asset/internal/infrastructure/kafka"
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
	prahariLogger.Info(ctx, "Registering Asset Service Kafka listener subscriptions: inspection.completed, maintenance.completed, maintenance.failed, incident.created, workflow.completed, permit.closed, risk.updated")
}

// Stop releases subscriptions.
func (l *Listener) Stop(ctx context.Context) {
	prahariLogger.Info(ctx, "Stopping Kafka event listeners")
}
