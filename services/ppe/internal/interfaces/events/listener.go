package events

import (
	"context"

	"prahari/services/ppe/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting PPE Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: permit.approved, visitor.checkedin, contractor.created, shift.started, risk.assessment.completed, workflow.completed")
	}()
}
