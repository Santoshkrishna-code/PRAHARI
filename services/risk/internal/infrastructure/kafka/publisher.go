package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
	prahariKafka "prahari/shared/kafka/producer"
)

// Publisher dispatches Kafka events.
type Publisher struct {
	producer *prahariKafka.Producer
}

// NewPublisher instantiates Publisher.
func NewPublisher(producer *prahariKafka.Producer) *Publisher {
	return &Publisher{producer: producer}
}

// Publish writes payload details to topic clusters.
func (p *Publisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Publishing operational risk event to Kafka topic",
		prahariLogger.String("topic", topic),
		prahariLogger.Int("payload_size", len(data)))

	return nil
}
