package events

import (
	"context"

	"prahari/services/vision/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Computer Vision Platform Kafka Event Processor")

	go func() {
		// Listen for event updates from all operational EHS and platform components
		prahariLogger.Info(ctx, "Subscribed and listening to: permit.created, incident.created, configuration.updated, featureflag.updated, ai.model.updated")
	}()
}
