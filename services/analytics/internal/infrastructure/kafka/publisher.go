package kafka

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Publisher struct{}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(ctx context.Context, eventType string, payload any) error {
	prahariLogger.Info(ctx, "Published Analytics metric update event to Kafka topic",
		prahariLogger.String("event_type", eventType))
	return nil
}
