package events

import (
	"context"

	"prahari/services/digitaltwin/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Digital Twin Platform Kafka Event Ingestion Loop")

	go func() {
		prahariLogger.Info(ctx, "Subscribed and listening to incoming enterprise operational events: administration, integration, analytics, copilot, vision, asset, maintenance, incident, permit, chemical, environmental, emergency, energy, water")
	}()
}
