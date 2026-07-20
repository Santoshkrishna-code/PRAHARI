package events

import (
	"context"

	"prahari/services/integration/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Enterprise Integration Hub Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: workflow.completed, tenant.created, configuration.updated, incident.created, maintenance.created, chemical.created")
	}()
}
