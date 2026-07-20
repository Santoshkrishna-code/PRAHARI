package events

import (
	"context"

	"prahari/services/loto/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting LOTO & Isolation Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: permit.created, maintenance.workorder.created, asset.updated, shift.started, workflow.completed")
	}()
}
