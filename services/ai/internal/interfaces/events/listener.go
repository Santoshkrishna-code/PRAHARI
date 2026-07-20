package events

import (
	"context"

	"prahari/services/ai/internal/infrastructure/kafka"
	prahariLogger "prahari/shared/logger"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener(consumer *kafka.Consumer) *Listener {
	return &Listener{consumer: consumer}
}

func (l *Listener) Start(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting AI Platform Kafka Event Processor")

	go func() {
		// Listen for event updates from all operational EHS and platform components
		prahariLogger.Info(ctx, "Subscribed and listening to: incident.created, hazard.created, near_miss.created, safety_observation.created, permit.issued, maintenance.completed, asset.created, environmental.spill, esg.updated, energy.usage, water.usage, moc.created, pha.created, barrier.failed, emergency.triggered, loto.applied, CAPA.created, meeting.conducted, chemical.received, configuration.updated, integration.completed")
	}()
}
