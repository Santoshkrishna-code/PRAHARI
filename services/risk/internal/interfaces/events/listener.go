package events

import (
	"context"

	prahariLogger "prahari/shared/logger"

	kafkaInfra "prahari/services/risk/internal/infrastructure/kafka"
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
	prahariLogger.Info(ctx, "Registering Operational Risk Kafka listener: hazard.created, hazard.updated, incident.created, incident.closed, nearmiss.created, observation.created, inspection.completed, maintenance.completed, permit.issued, permit.closed, workflow.completed")
}

// Stop releases subscriptions.
func (l *Listener) Stop(ctx context.Context) {
	prahariLogger.Info(ctx, "Stopping Kafka event listeners")
}
