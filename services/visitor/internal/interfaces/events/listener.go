package events

import (
	"context"

	"prahari/services/visitor/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Visitor Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: contractor.created, permit.approved, emergency.declared, shift.started, workflow.completed")
	}()
}
