package events

import (
	"context"

	"prahari/services/shift/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Shift Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: permit.approved, maintenance.workorder.created, incident.created, emergency.declared, workflow.completed")
	}()
}
