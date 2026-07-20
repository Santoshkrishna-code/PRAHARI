package events

import (
	"context"

	"prahari/services/administration/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Enterprise Administration Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: identity.created, workflow.completed, notification.sent")
	}()
}
