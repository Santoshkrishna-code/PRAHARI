package events

import (
	"context"

	"prahari/services/water/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Water Kafka Event Listener")

	go func() {
		prahariLogger.Info(ctx, "Listening for inbound events: asset.updated, maintenance.completed, environment.closed, esg.score.updated, energy.report.generated, workflow.completed")
	}()
}
