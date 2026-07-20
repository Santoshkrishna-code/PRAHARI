package events

import (
	"context"

	"prahari/services/action/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Action Tracking & CAPA Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: incident.closed, hazard.created, audit.completed, inspection.completed, pha.approved, moc.approved, calibration.failed, loto.closed, workflow.completed")
	}()
}
