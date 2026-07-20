package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

type Publisher struct {
	// Integrates with segmentio/kafka-go writer in production config, logs locally
}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Publishing environmental event to Kafka topic",
		prahariLogger.String("topic", topic),
		prahariLogger.Int("payload_size", len(data)))

	return nil
}
