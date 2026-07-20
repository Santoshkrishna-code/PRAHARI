package events

import (
	"context"

	"prahari/services/document/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Document Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: moc.approved, pha.approved, audit.completed, workflow.completed")
	}()
}
