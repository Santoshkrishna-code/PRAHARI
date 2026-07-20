package events

import (
	"context"

	"prahari/services/bcm/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Business Continuity Management (BCM) Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: emergency.declared, incident.created, asset.updated, risk.assessment.completed, workflow.completed")
	}()
}
