package events

import (
	"context"

	"prahari/services/calibration/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Calibration Management Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: asset.created, maintenance.workorder.created, inspection.started, permit.created, workflow.completed")
	}()
}
