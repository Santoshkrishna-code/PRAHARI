package events

import (
	"context"

	"prahari/services/chemical/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Chemical Safety & SDS Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: permit.created, incident.created, maintenance.started, environmental.alert, workflow.completed, meeting.completed, action.closed")
	}()
}
