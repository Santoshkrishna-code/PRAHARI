package events

import (
	"context"

	"prahari/services/meetings/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Meetings & Toolbox Talks Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: shift.started, permit.created, incident.closed, hazard.created, workflow.completed")
	}()
}
