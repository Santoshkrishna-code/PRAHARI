package events

import (
	"context"

	"prahari/services/environmental/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Environmental Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: incident.created, maintenance.completed, risk.activated, audit.closed, compliance.reviewed, workflow.completed")
	}()
}
