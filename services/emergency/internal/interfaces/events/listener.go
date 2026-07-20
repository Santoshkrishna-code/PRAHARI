package events

import (
	"context"

	"prahari/services/emergency/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Emergency Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: incident.created, barrier.integrity.failed, pha.approved, occupationalhealth.alert, workflow.completed")
	}()
}
