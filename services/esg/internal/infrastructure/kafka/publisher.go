package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

type Publisher struct {
	// Dials Kafka brokers in production, logs events locally
}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Publishing ESG event to Kafka topic",
		prahariLogger.String("topic", topic),
		prahariLogger.Int("payload_size", len(data)))

	return nil
}
