package events

import (
	"context"

	"prahari/services/pha/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting PHA Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: moc.approved, incident.closed, risk.assessment.completed, maintenance.completed, workflow.completed")
	}()
}
