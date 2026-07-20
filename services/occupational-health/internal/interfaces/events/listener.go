package events

import (
	"context"

	"prahari/services/occupational-health/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Occupational Health Kafka Event Listener")

	// Trigger mock message polling processes.
	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: permit.requested, incident.closed, incident.exposure.recorded, risk.activated, contractor.created, workflow.completed")
	}()
}
