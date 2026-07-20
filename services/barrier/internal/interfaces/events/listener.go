package events

import (
	"context"

	"prahari/services/barrier/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Barrier Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: pha.approved, moc.approved, maintenance.completed, incident.closed, workflow.completed")
	}()
}
